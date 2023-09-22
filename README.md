# Shrink URL

## Notes

1. The sample code was in C#, a language that I have never used, so I tried to rewrite it and implement the API in Golang.
2. I am not sure what should I pass as input to the encryptor, which I assumed to be the full URL, but I was unable to generate the expected output even from an exact clone of the C# program. The output from the Golang cryptography class is the exact match with the output from the C# sample.
3. Logging is not implemented.

## Instructions

### Start dev server

Default port: 80

```bash
go run ./cmd/app --port=<custom port>
```

### Run tests

```bash
go test ./test
```

### Build (for Linux)

```bash
make build
```

### Run server from binary

```bash
./bin/shrink-url --port=<custom port>
```

### Build and run on Docker

Default expose port: 4000 (update Makefile to use a different port)

```bash
make up_build
```

## API

<details>
 <summary><code>GET</code> <code><b>/</b></code></summary>

##### Query

> | key   | type     | description    |
> | ----- | -------- | -------------- |
> | `url` | `string` | `original URL` |

##### Responses

> | http code | content-type       | response                                |
> | --------- | ------------------ | --------------------------------------- |
> | `200`     | `application/text` | `<base URL>/<cipher from original URL>` |
> | `400`     | `application/text` | `<error message>`                       |
> | `404`     | `application/text` | `404 page not found`                    |

</details>

## Improvements

1. Encryption key should be loaded as environment variable or stored in a separated file instead of hardcoded in the program. Depending on the security level, the salt might have to be managed in the same way as the encryption key.
2. Although it is the requirement of the task, I do not see the need of using encryption over a simple hash function, unless the original URL has to be recovered at some point. However, if that is the case, the original URL might still not be recoverable as the output might have been trimmed.
3. Output generated from the provided cryptography class always start with the same bytes when the input fit within the same block numbers. Mainly because a full URL mostly starts with "http://" or "https://". If the MaxLength is too small, output might always be the same. I am not sure whether there were mistakes on the implementation of the Rijndael algorithm.
4. I do not understand the real usage of this program, thus, it is difficult to make any suggestions. If the sole purpose is to generate a short URL that does not have to be recoverable, I would suggest using a hash function like SHA256, instead of encryption. That would make the implementation simplier, remove need of using and storing encryption keys, as well as ensuring uniqueness of output.
