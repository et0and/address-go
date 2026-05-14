# address

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/et0and/address-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Challenge

Response Types:

- <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#ChallengeGetResponse">ChallengeGetResponse</a>

Methods:

- <code title="get /challenge">client.Challenge.<a href="https://pkg.go.dev/github.com/et0and/address-go#ChallengeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#ChallengeGetResponse">ChallengeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RequestKey

Response Types:

- <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#RequestKeyNewResponse">RequestKeyNewResponse</a>

Methods:

- <code title="post /request-key">client.RequestKey.<a href="https://pkg.go.dev/github.com/et0and/address-go#RequestKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#RequestKeyNewParams">RequestKeyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#RequestKeyNewResponse">RequestKeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Addresses

Response Types:

- <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#AddressGetResponse">AddressGetResponse</a>
- <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#AddressListResponse">AddressListResponse</a>

Methods:

- <code title="get /v1/addresses/{id}">client.Addresses.<a href="https://pkg.go.dev/github.com/et0and/address-go#AddressService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#AddressGetResponse">AddressGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/addresses">client.Addresses.<a href="https://pkg.go.dev/github.com/et0and/address-go#AddressService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#AddressListParams">AddressListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#AddressListResponse">AddressListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Search

Response Types:

- <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#SearchQueryResponse">SearchQueryResponse</a>

Methods:

- <code title="get /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/et0and/address-go#SearchService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#SearchQueryParams">SearchQueryParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#SearchQueryResponse">SearchQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Reverse

Response Types:

- <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#ReverseGeocodeResponse">ReverseGeocodeResponse</a>

Methods:

- <code title="get /v1/reverse">client.Reverse.<a href="https://pkg.go.dev/github.com/et0and/address-go#ReverseService.Geocode">Geocode</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#ReverseGeocodeParams">ReverseGeocodeParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#ReverseGeocodeResponse">ReverseGeocodeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Meta

Response Types:

- <a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#MetaGetResponse">MetaGetResponse</a>

Methods:

- <code title="get /v1/meta">client.Meta.<a href="https://pkg.go.dev/github.com/et0and/address-go#MetaService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/et0and/address-go">address</a>.<a href="https://pkg.go.dev/github.com/et0and/address-go#MetaGetResponse">MetaGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
