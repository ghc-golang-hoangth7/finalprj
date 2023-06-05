├── grpc
    ├── flights
        ├──...
        ├──server.go                        # main file of flights grpc service, contain startup, load config, etc
        ├──.env                             # config of flights service
    ├── planes
        ├──...
        ├──server.go                        # main file of planes grpc service, contain startup, load config, etc
        ├──.env                             # config of planes service

├──db                                       # Generated grpc files
    ├──db.sql                               # Generated grpc files

├──pb                                       # Generated grpc files
├──proto                                    # Protobuff files
├──clients                                  # API endpoint
    ├──graph                                # GraphQL endpoint
        ├──cmd
            ├──server.go                    # main file of GraphQL endpoint, contain startup, load config, etc

.
├── docker-compose.yml
├── db
│   ├── db.sql
│   └── Dockerfile
├── grpc
│   ├── flights
│   │    ├── cmd
│   │    │   ├── server.go
│   │    │   └── Dockerfile
│   │    └── ...<files require by flights service>
│   └── planes
│       ├── cmd
│       │   ├── server.go
│       │   └── Dockerfile
│       └── ...<files require by planes service>
├── models
│   └── ...<files require by flights/planes service and graphql endpoint>
├── pb
│   └── ...<files require by flights/planes service and graphql endpoint>
├── proto
│   └── ...<file require by flights/planes service and graphql endpoint>
├── common
│   └── ...<file require by flights/planes service and graphql endpoint>
├── clients
│   └── graph
│       ├── cmd
│       │   ├── server.go
│       │   └── Dockerfile
│       └── ...<file require by graphql endpoint>
│
├── go.mod      #file require by flights/planes service and graphql endpoint
└── go.sum      #file require by flights/planes service and graphql endpoint