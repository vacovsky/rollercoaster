---
draft: false
date: "2017-10-06T11:03:36"
tags: ["LoggingDB",
"ServiceA-DB",
"IdentityProvider",
]
title: Service-A
categories: ["databases-and-clients"]
depmap: [ "graph LR",
"style Service-A fill:#dc6a12,stroke:#8c2231,stroke-width:2px",
"Service-A -->LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-A -->ServiceA-DB(\"fa:fa-database ServiceA-DB\")",
"Service-A -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
]
---
