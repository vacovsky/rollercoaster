---
draft: false
date: "2017-10-02T15:03:19"
tags: ["Crypto",
"MessageBus-A",
"File_Share",
"LoggingDB",
"IdentityDB",
]
title: IdentityProvider
categories: ["everything"]
depmap: [ "graph LR",
"style IdentityProvider fill:#ceb627,stroke:#482954,stroke-width:2px",
"IdentityProvider -->|HTTP|Crypto((\"fa:fa-globe Crypto\"))",
"IdentityProvider -->|WCF|MessageBus-A{\"fa:fa-tasks MessageBus-A\"}",
"IdentityProvider -->|SMB/CIFS/NFS|File_Share[\"fa:fa-files-o File_Share\"]",
"IdentityProvider -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"IdentityProvider -->|SQL|IdentityDB(\"fa:fa-database IdentityDB\")",
]
---
			