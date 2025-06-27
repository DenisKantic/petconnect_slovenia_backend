### HTTP status code summary

<br>

| STATUS | MEANING           | DESCRIPTION                                                                     |
|--------|-------------------|---------------------------------------------------------------------------------|
| 200    | OK                | Everything worked as expected                                                   |
| 400    | Bad Request       | The request failed due to missing some required parameter                       |
| 401    | Unauthorized      | No valid API key provided                                                       |
| 402    | Request Failed    | The parameters were valid but the request failed                                |
| 403    | Forbidden         | The API key doesn't have permissions to perform the request                     |
| 404    | Not Found         | The requested source doesn't exist                                              |
| 429    | Too Many Requests | Too many requests hit the API quickly. Rate limiter blocking this               |
| 500    | Server Error      | Some action cannot be processed on the server, check the log for detailed error |





