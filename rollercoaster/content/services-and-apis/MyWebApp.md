---
draft: false
date: "2017-10-02T15:03:19"
tags: ["ServiceAPI",
"IdentityProvider",
"Service-A",
]
title: MyWebApp
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style MyWebApp fill:#ae0f99,stroke:#229555,stroke-width:2px",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|ServiceAPI((\"fa:fa-globe ServiceAPI\"))",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|WCF|Service-A{\"fa:fa-tasks Service-A\"}",
]
---
			