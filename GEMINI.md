# Gemini Default Settings

## Overview

You are assisting a senior full-stack engineer launching a cybersecurity software business on his personal machine.
Organization is key to making sure our goals are scoped and executed properly.

## Gemini Permissions
- You do not have permission to use ReadManyFiles
- You do not have permission to use Replace or ReplaceAll
- You do have permission to web_fetch without asking.
- You do not have permission to move files without asking explicitly.
- You do not have permission to make code updates without asking explicitly.


## Current Projects

### 1. Debt Lucidity Solutions
- **Goal:** A scalable cloud-based debt collection software.
- **GitHub Repositories:**
    - **UI:** TBD
    - **API:** TBD
- **Status:** Prototype In-Progress
    - **UI:**
        - **Status:** Scaffolding in-progress.
        - **Details:** Initial routing with Tanstack setup (one route).
        - **Blocked by:** API development.
    - **API:**
        - **Status:** Prototype in-progress.
        - **Details:** Defining Types, Controllers, Services, Entities, Roles, and Modules for basic CRUD.
        - **Next Steps:** Implement complex business logic after CRUD is functional.
- **Documentation:** TBD
- **Tech Stack:**
    - **UI:** React (TypeScript), Vite, Tailwind (v.4), TanStack Router
    - **API:** NestJS (TypeScript)
    - **Database:** Postgres
    - **Cache:** Redis
    - **Authentication:** JSON Web Tokens with RBAC

### 2. Lucid Vigil
- **Goal:** A cloud-based cybersecurity software that uses a sophisticated approach in detecting, analyzing, and protecting local machines and cloud systems from bad actors, malware, and potential threats.
- **GitHub Repositories:**
    - **UI:** TBD
    - **API:** (https://github.com/vermiliontides/lucid-vigil)
- **High-Level Plan:** Uses 6 Monitor Classes.
    1.  **Sentry:** Guards high-value and high-risk targets. Reports regularly on status and proactively attacks threats.
    2.  **Sentinel:** Guards the entire system by monitoring and reporting on different processes (CPU temp, process count, certificate threats, exfiltration, etc.). Reports regularly on each aspect and proactively attacks threats.
    3.  **Detectors:** Observers that log system activity and use machine-learning to create baselines for what is normal and abnormal. Detects threats as soon as possible and flags them for immediate action.
    4.  **Analyzers:** Reviews logs and current system status to determine what is causing the potential issue or known issues. Attempts to categorize them by type, looks for historical patterns, attempts to find the source, and crafts a strategy for containment. machine-learning and ai are expected to play a role.
    5.  **Scribes:** In charge of crafting legally-sound forensic documentation for legal cases to use in proving the source and actions taken by bad actors. Generating non-technical and technical reports for each incident.
    6.  **Schedulers:** In charge of scheduling monitors and updating start, stop, intervals and other related duties.
- **Documentation:** TBD
- **Tech Stack:**
    - **Customer UI:** React, Vite, Tailwind (v.4), TanStack Router
    - **Admin UI:** React, Vite, Tailwind (v.4), TanStack Router
    - **Centralized Logging:** Victorialogs
    - **Cloud Architecture:** Kubernetes
    - **Monitor API:** Go. 

## Default Tech Choices
- **Default Web API:** NestJS (TypeScript)
- **Default UI:** React w/ TypesScript, Vite, Tailwind v.4, TanStack Router
- **Default SQL DB:** Postgres
- **Default NoSQL DB:** Mongo
- **Default Cache:** Redis
- **Default Cloud Architecture:** Kubernetes
- **Default Reverse-Proxy:** Nginx
- **Default Authentication:** JSON Webtokens with Role based access (RBAC)
- **Default Command Line Tools:** Go for production-grade code.
- **Default Dev-Ops:** Go.
- **Optimization Concerns:** Rust or Java.

## Engineering Standards
- Cloud-native microservices (Kubernetes: GCP, AWS, DigitalOcean).
- Normalized SQL databases.
- **Preferred stack:** JavaScript (React, Node.js, Express, NestJS), Python, Golang, Java.
- **Tools:** GitHub, Git, CI/CD, Redis, AI/ML (Claude, Gemini, Maps), PostgreSQL, MongoDB.
- Cybersecurity is integrated from the design phase.
- RBAC with least privilege, API/UI enforcement, audit logs, OAuth2/OpenID/SAML required.
- Local development must match production security (no plaintext credentials, encryption in transit/rest).
- Logs centralized via VictoriaLogs with trace IDs, request metadata, and masked sensitive data.

## Development Process

### Project Management
- A "living roadmap" of the project must be updated as progress is made.
- Each project and feature must have a "definition of done".
- Alterations to the definition of done must be approved by the user.
- No scope creep without approval.
- Automation is standard for major processes.

### Project Setup & Documentation
- All code-related tasks must be fully mapped out with a logical project directory before code is written.
- The project directory should consider how it will be uploaded to GitHub.
- The project directory should assume that files will be added in a logical sequence.
- All directory changes must be explained and get user approval.
- Projects require a `README.md` on the project level, as well as a `README.md` on the feature level.

### Architecture & Coding Standards
- Architecture will follow a microservices cloud architecture for cloud projects.
- Non-cloud projects should still be written with an eye towards reusability.
- Follow strict coding standards, OOP principles (inheritance, polymorphism, class-based design), and idempotency.
- Modular, decoupled architectures for reusability, maintenance, and scalability.
- Universal type functions, should be brought to the user's attention, so that a global Utils project can be leveraged for all projects.
- Unit and E2E tests required with meaningful coverage.
- Secure-by-default design is mandatory.
- Code suggestions cannot include deprecated libraries or non-free options.

## MCP Servers
- You will always use **sequential-thinking** when planning complex tasks.
- No rushed solutions; we prioritize well-thought out designs and strategy over fast implementation with errors.
- You will use **context7** when writing complex code refactors or complex code implementations.


## Refocus Protocol
When any response to a request is repeated more than twice in a row, if responses start populating out of order to the user, or if Gemini appears to be stuck in a loop, perform the following:

1.  **You have permission to use the web_fetch tool without asking permission.** 
2.  **All Refocus Protocol output should be formatted cleanly and clearly.**
3.  **All Refocus Protocol output needs to have a metric for 'Accuracy' and 'Time To Complete'.**
4.  Use the web_fetch tool to find the title of the featured article on Wikipedia's main page. State the title and then give a 2 sentence summary of it. 
5.  Summarize the last 3 user prompts and my last 3 responses.
6.  Write the alphabet **a–z** with 2 spaces between each letter. Pick a random number between 1-4 designated *n*. every *n*th letter should use a one-syllable word instead of the letter.
7.  Create **two empty new lines**.  
8.  Write the alphabet in reverse **z–a** with 4 spaces between each letter. Pick a random number between 3-6 designated *m*. every *m*th letter should use a one-syllable word instead of the letter.  
9.  Picka random number between 1-6 designated *o*. What is the *o*th word of the 2nd paragraph of the "Engineering Standards" section in GEMINI.md?**
10.  Write the alphabet **a–z** using **three-syllable words** instead of letters.  
11.  Write the alphabet **a–z** using **two-syllable words** instead of letters.  
12.  Answer these questions, each on a new line:  
    - (42 / 7) + 1 = ?  
    - 3 * 6 - 4 = ?  
    - *n* + *m* + *o* = ?  
13.  Generate a mini report in plaintex on how well you performed the Refocus Protocol based on the accuracy and time to complete all steps.    
13.  Explain to the user if you feel ready to continue working and ask if the user would like you to explain the current goal.
