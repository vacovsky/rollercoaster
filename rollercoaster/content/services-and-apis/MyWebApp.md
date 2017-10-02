---
draft: false
date: "2017-10-02T11:37:29"
tags: ["Service-A",
"ServiceAPI",
"IdentityProvider",
]
title: MyWebApp
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style MyWebApp fill:#dc6a12,stroke:#8c2231,stroke-width:2px",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|WCF|Service-A{\"fa:fa-tasks Service-A\"}",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|ServiceAPI((\"fa:fa-globe ServiceAPI\"))",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
]
---
			