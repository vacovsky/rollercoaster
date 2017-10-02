---
draft: false
date: "2017-10-02T14:57:28"
tags: ["File_Share",
"LoggingDB",
"IdentityDB",
"Crypto",
"MessageBus-A",
]
title: IdentityProvider
categories: ["everything"]
depmap: [ "graph LR",
"style IdentityProvider fill:#b36431,stroke:#277d78,stroke-width:2px",
"IdentityProvider -->|SMB/CIFS/NFS|File_Share[\"fa:fa-files-o File_Share\"]",
"IdentityProvider -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"IdentityProvider -->|SQL|IdentityDB(\"fa:fa-database IdentityDB\")",
"IdentityProvider -->|HTTP|Crypto((\"fa:fa-globe Crypto\"))",
"IdentityProvider -->|WCF|MessageBus-A{\"fa:fa-tasks MessageBus-A\"}",
]
---
			