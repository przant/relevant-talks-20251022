# XML to JSON Conversion: Go vs Dynamic Languages

## The Problem

Convert SOAP XML envelope to custom JSON payload for downstream processing.

## The Python/Lambda Approach (4 Steps, 4 Functions)
```python
# Lambda 1: Parse SOAP envelope
def parse_soap(event):
    soap_xml = event['body']
    parsed = xml.etree.ElementTree.fromstring(soap_xml)
    return {
        'parsed_envelope': extract_envelope_data(parsed)
    }

# Lambda 2: Adapter - extract relevant fields
def adapter(event):
    envelope = event['parsed_envelope']
    relevant_data = {
        'field1': envelope['SoapBody']['Field1'],
        'field2': envelope['SoapBody']['Field2'],
        # ... manual field extraction
    }
    return {'adapted_data': relevant_data}

# Lambda 3: Transform to custom JSON
def transformer(event):
    adapted = event['adapted_data']
    custom_json = {
        'customField1': adapted['field1'],
        'customField2': adapted['field2'],
        # ... manual field mapping
    }
    return {'payload': custom_json}

# Lambda 4: Do actual work
def worker(event):
    payload = event['payload']
    # Finally do the real work
    process_data(payload)
```

**Cost:**
- 4 Lambda invocations
- 4 separate deployments
- Inter-Lambda communication overhead
- Complex error handling across boundaries
- Difficult to debug and trace

---

## The Go Approach (1 Function, 3 Lines)
```go
// Define what you need (once)
type SoapEnvelope struct {
    Field1 string `xml:"Body>Field1" json:"customField1"`
    Field2 string `xml:"Body>Field2" json:"customField2"`
}

// Single function
func handler(w http.ResponseWriter, r *http.Request) {
    var envelope SoapEnvelope
    
    // Step 1: Decode XML
    xml.NewDecoder(r.Body).Decode(&envelope)
    
    // Step 2: Encode JSON (struct is the bridge)
    json.NewEncoder(w).Encode(envelope)
    
    // Step 3: Do actual work
    processData(envelope)
}
```

**Benefits:**
- Single deployment unit
- Type-safe at compile time
- No runtime field mapping errors
- Clear data flow
- Easy to test and debug
- Zero inter-service communication overhead

---

## Why This Works in Go

### Struct Tags Handle Everything
```go
type Data struct {
    XMLName  xml.Name `xml:"SoapEnvelope"`
    Field1   string   `xml:"Body>Field1" json:"customField1"`
    Ignored  string   `xml:"-" json:"-"`  // Skip unwanted fields
}
```

- `xml:"Body>Field1"` - XPath-like navigation
- `json:"customField1"` - Custom JSON field names
- `xml:"-"` - Ignore fields you don't need
- `omitempty` - Skip zero values

### The Struct Is The Bridge
```
XML bytes → xml.Decode() → Struct → json.Encode() → JSON bytes
```

No loops. No manual field copying. No reflection magic you can't debug.

### Type Safety Catches Errors Early
```python
# Python - Runtime error in production
payload['cutsomField1']  # Typo, crashes at runtime
```
```go
// Go - Compile error before deployment
envelope.CutsomField1  // Won't compile
```

---

## Real-World Comparison

| Aspect | Python/Lambda (4-step) | Go (single function) |
|--------|------------------------|----------------------|
| **Lines of code** | ~100+ lines | ~10 lines |
| **AWS resources** | 4 Lambdas + SQS/SNS | 1 Lambda or Container |
| **Monthly cost** | 4× invocation cost | 1× invocation cost |
| **Debugging** | Trace across 4 services | Single stack trace |
| **Type safety** | Runtime only | Compile time |
| **Field mapping errors** | Discovered in production | Caught before deploy |
| **Cold start penalty** | 4× cold starts possible | 1× cold start |

---

## When Dynamic Languages Make Sense

- Rapid prototyping without known schema
- Data exploration and experimentation
- Schemas change frequently with no control
- Small scripts and one-off tasks

## When Go Makes Sense

- Production systems with known schemas
- Performance-critical parsing
- Type safety requirements
- Clear data contracts between services
- Long-running services (not serverless)

---

## The Bottom Line

**Python approach:** Flexible, but pays cost in complexity and runtime errors

**Go approach:** Rigid structure upfront, but eliminates entire classes of errors and architectural complexity

For production data pipelines with stable schemas, Go's "struct as bridge" pattern eliminates unnecessary architectural layers while providing compile-time guarantees.