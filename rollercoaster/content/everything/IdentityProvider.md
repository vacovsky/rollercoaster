---
draft: false
date: "2017-10-06T12:02:14"
tags: ["LoggingDB",
"IdentityDB",
"Crypto",
"MessageBus-A",
"File_Share",
]
title: IdentityProvider
categories: ["everything"]
depmap: [ "graph LR",
"style IdentityProvider fill:#2943b7,stroke:#4e2d62,stroke-width:2px",
"IdentityProvider -->LoggingDB(\"fa:fa-database LoggingDB\")",
"IdentityProvider -->IdentityDB(\"fa:fa-database IdentityDB\")",
"IdentityProvider -->Crypto((\"fa:fa-globe Crypto\"))",
"IdentityProvider -->MessageBus-A[\"fa:fa-sitemap MessageBus-A\"]",
"IdentityProvider -->File_Share[\"fa:fa-files-o File_Share\"]",
]
---
