# Jaeger

To enable the Jaeger tracer:

```yaml tab="File (YAML)"
tracing:
  jaeger: {}
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger]
```

```bash tab="CLI"
--tracing.jaeger=true
```

!!! warning
    Traefik is able to send data over the compact thrift protocol to the [Jaeger agent](https://www.jaegertracing.io/docs/deployment/#agent)
    or a [Jaeger collector](https://www.jaegertracing.io/docs/deployment/#collector).

!!! info
    All Jaeger configuration can be overridden by [environment variables](https://github.com/jaegertracing/jaeger-client-go#environment-variables)

#### `samplingServerURL`

_Required, Default="http://localhost:5778/sampling"_

Address of the Jaeger Agent HTTP sampling server.

```yaml tab="File (YAML)"
tracing:
  jaeger:
    samplingServerURL: http://localhost:5778/sampling
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger]
    samplingServerURL = "http://localhost:5778/sampling"
```

```bash tab="CLI"
--tracing.jaeger.samplingServerURL=http://localhost:5778/sampling
```

#### `samplingType`

_Required, Default="const"_

Type of the sampler.

Valid values are:

- `const`
- `probabilistic`
- `rateLimiting`

```yaml tab="File (YAML)"
tracing:
  jaeger:
    samplingType: const
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger]
    samplingType = "const"
```

```bash tab="CLI"
--tracing.jaeger.samplingType=const
```

#### `samplingParam`

_Required, Default=1.0_

Value passed to the sampler.

Valid values are:

- for `const` sampler, 0 or 1 for always false/true respectively
- for `probabilistic` sampler, a probability between 0 and 1
- for `rateLimiting` sampler, the number of spans per second

```yaml tab="File (YAML)"
tracing:
  jaeger:
    samplingParam: 1.0
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger]
    samplingParam = 1.0
```

```bash tab="CLI"
--tracing.jaeger.samplingParam=1.0
```

#### `localAgentHostPort`

_Required, Default="127.0.0.1:6831"_

Local Agent Host Port instructs the reporter to send spans to the Jaeger Agent at this address (host:port).

```yaml tab="File (YAML)"
tracing:
  jaeger:
    localAgentHostPort: 127.0.0.1:6831
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger]
    localAgentHostPort = "127.0.0.1:6831"
```

```bash tab="CLI"
--tracing.jaeger.localAgentHostPort=127.0.0.1:6831
```

#### `gen128Bit`

_Optional, Default=false_

Generates 128 bits trace IDs, compatible with OpenCensus.

```yaml tab="File (YAML)"
tracing:
  jaeger:
    gen128Bit: true
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger]
    gen128Bit = true
```

```bash tab="CLI"
--tracing.jaeger.gen128Bit
```

#### `propagation`

_Required, Default="jaeger"_

Sets the propagation header type.

Valid values are:

- `jaeger`, jaeger's default trace header.
- `b3`, compatible with OpenZipkin

```yaml tab="File (YAML)"
tracing:
  jaeger:
    propagation: jaeger
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger]
    propagation = "jaeger"
```

```bash tab="CLI"
--tracing.jaeger.propagation=jaeger
```

#### `traceContextHeaderName`

_Required, Default="uber-trace-id"_

HTTP header name used to propagate tracing context.
This must be in lower-case to avoid mismatches when decoding incoming headers.

```yaml tab="File (YAML)"
tracing:
  jaeger:
    traceContextHeaderName: uber-trace-id
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger]
    traceContextHeaderName = "uber-trace-id"
```

```bash tab="CLI"
--tracing.jaeger.traceContextHeaderName=uber-trace-id
```

### disableAttemptReconnecting

_Optional, Default=true_

Disables the UDP connection helper that periodically re-resolves the agent's hostname and reconnects if there was a change.
Enabling the re-resolving of UDP address make the client more robust in Kubernetes deployments.

```yaml tab="File (YAML)"
tracing:
  jaeger:
    disableAttemptReconnecting: false
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger]
    disableAttemptReconnecting = false
```

```bash tab="CLI"
--tracing.jaeger.disableAttemptReconnecting=false
```

### `collector`
#### `endpoint`

_Optional, Default=""_

Collector Endpoint instructs the reporter to send spans to the Jaeger Collector at this URL.

```yaml tab="File (YAML)"
tracing:
  jaeger:
    collector:
        endpoint: http://127.0.0.1:14268/api/traces?format=jaeger.thrift
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger.collector]
    endpoint = "http://127.0.0.1:14268/api/traces?format=jaeger.thrift"
```

```bash tab="CLI"
--tracing.jaeger.collector.endpoint=http://127.0.0.1:14268/api/traces?format=jaeger.thrift
```

#### `user`

_Optional, Default=""_

User instructs the reporter to include a user for basic HTTP authentication when sending spans to the Jaeger Collector.

```yaml tab="File (YAML)"
tracing:
  jaeger:
    collector:
        user: my-user
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger.collector]
    user = "my-user"
```

```bash tab="CLI"
--tracing.jaeger.collector.user=my-user
```

#### `password`

_Optional, Default=""_

Password instructs the reporter to include a password for basic HTTP authentication when sending spans to the Jaeger Collector.

```yaml tab="File (YAML)"
tracing:
  jaeger:
    collector:
        password: my-password
```

```toml tab="File (TOML)"
[tracing]
  [tracing.jaeger.collector]
    password = "my-password"
```

```bash tab="CLI"
--tracing.jaeger.collector.password=my-password
```
