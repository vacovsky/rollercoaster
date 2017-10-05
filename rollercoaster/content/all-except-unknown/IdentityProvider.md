---
draft: false
date: "2017-10-05T15:23:02"
tags: ["LoggingDB",
"IdentityDB",
"Crypto",
"MessageBus-A",
"File_Share",
]
title: IdentityProvider
categories: ["all-except-unknown"]
depmap: [ "graph LR",
"style IdentityProvider fill:#d67f26,stroke:#505d2b,stroke-width:2px",
"IdentityProvider -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"IdentityProvider -->|SQL|IdentityDB(\"fa:fa-database IdentityDB\")",
"IdentityProvider -->|HTTP|Crypto((\"fa:fa-globe Crypto\"))",
"IdentityProvider -->|AMQP|MessageBus-A[\"fa:fa-sitemap MessageBus-A\"]",
"IdentityProvider -->|SMB/CIFS/NFS|File_Share[\"fa:fa-files-o File_Share\"]",
]
---
			