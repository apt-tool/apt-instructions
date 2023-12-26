package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

var queries = []string{
	`query IntrospectionQuery {
  __schema {
    queryType {
      name
    }
    mutationType {
      name
    }
    subscriptionType {
      name
    }
    types {
      ...FullType
    }
    directives {
      name
      description
      locations
      args {
        ...InputValue
      }
    }
  }
}
fragment FullType on __Type {
  kind
  name
  description
  fields(includeDeprecated: true) {
    name
    description
    args {
      ...InputValue
    }
    type {
      ...TypeRef
    }
    isDeprecated
    deprecationReason
  }
  inputFields {
    ...InputValue
  }
  interfaces {
    ...TypeRef
  }
  enumValues(includeDeprecated: true) {
    name
    description
    isDeprecated
    deprecationReason
  }
  possibleTypes {
    ...TypeRef
  }
}
fragment InputValue on __InputValue {
  name
  description
  type {
    ...TypeRef
  }
  defaultValue
}
fragment TypeRef on __Type {
  kind
  name
  ofType {
    kind
    name
    ofType {
      kind
      name
      ofType {
        kind
        name
        ofType {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
              }
            }
          }
        }
      }
    }
  }
}`,
	`query brokenAccessControl {
  myInfo(accessToken:"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJwb2MiLCJzdWIiOiJKdWxpZW4iLCJpc3MiOiJBdXRoU3lzdGVtIiwiZXhwIjoxNjAzMjkxMDE2fQ.r3r0hRX_t7YLiZ2c2NronQ0eJp8fSs-sOUpLyK844ew", veterinaryId: 2){
    id, name, dogs {
      name
    }
  }
}`,
	`query sqli {
  dogs(namePrefix: "ab%' UNION ALL SELECT 50 AS ID, C.CFGVALUE AS NAME, NULL AS VETERINARY_ID FROM CONFIG C LIMIT ? -- ", limit: 1000) {
    id
    name
  }
}`,
	`query xss  {
  myInfo(veterinaryId:"<script>alert('1')</script>" ,accessToken:"<script>alert('1')</script>") {
    id
    name
  }
}`,
	`query dos {
  allDogs(onlyFree: false, limit: 1000000) {
    id
    name
    veterinary {
      id
      name
      dogs {
        id
        name
        veterinary {
          id
          name
          dogs {
            id
            name
            veterinary {
              id
              name
              dogs {
                id
                name
                veterinary {
                  id
                  name
                  dogs {
                    id
                    name
                    veterinary {
                      id
                      name
                      dogs {
                        id
                        name
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}`,
	`query {
  Veterinary(id: "1") {
    name
  }
  second:Veterinary(id: "2") {
    name
  }
  third:Veterinary(id: "3") {
    name
  }
}`,
	`query {
  auth(veterinaryName: "Julien")
  second: auth(veterinaryName:"Benoit")
}`,
}

func send(host string, query string) bool {
	ctx := context.Background()
	client := graphql.NewClient(host)

	// make a request
	req := graphql.NewRequest(query)

	// set any variables
	req.Var("key", "value")

	// run it and capture the response
	var respData json.Decoder
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Println(err)

		return false
	}

	return true
}

func main() {
	var (
		hostFlag = flag.String("host", "localhost", "target host address")
	)

	flag.Parse()

	for _, query := range queries {
		if send(*hostFlag, query) {
			os.Exit(1)
		}
	}

	os.Exit(0)
}
