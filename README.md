# test_task_trood_ai_helpdesk
Test task for golang developer

## Description
```
The project is an example of a simple system for automated customer support. The system receives a request and sends it to a RabbitMQ queue. Workers process the received message. It is passed to an NLP model, and based on the extracted intents, a response is returned from db. If no response is found, the message is forwarded to a human support agent, the functionality of which is implemented through the terminal. The support agent's response is then returned to the client.
```

## How to use
### in first terminal (nlp server) :
* pip3 install flask
* pip3 install spacy
* python3 -m spacy download en_core_web_sm
* python3 nlp_server/nlp_server.py

### in second terminal (main api) : 
* docker-compose up -d
* go run cmd/main.go

### in third terminal (workers) :
* go run worker/main.go

### in the fourth terminal (helpdesk):
here you can answer questions that were not found in the database
* go run human/main.go 

for test you can use :
```
curl "http://localhost:3000/helpdesk?msg=i%20want%20to%20issue%20refund"
```


## Task

```
Backend Engineer Task (Brief Solution Proposal + Code)

Scenario:
You are tasked with designing a back-end system to automate customer support using AI agents. The system should handle customer queries in real time, use NLP to understand them, and respond with appropriate answers from a knowledge base or escalate to a human agent if needed.

Key Requirements:
Use AI/NLP for query understanding.
Connect to a database to retrieve answers.
Ensure the system can scale efficiently.

Task:
Design a high-level architecture for the back-end system:
Briefly describe the components of the system, such as:
The AI agent for understanding customer queries.
A database to store knowledge and provide responses.
A queueing system for managing real-time interactions.
Ensure scalability to handle multiple requests concurrently.

Write a piece of code to:
Simulate a simple query processing system.
Retrieve a response from a database (can be a mock or in-memory DB).
Use a basic NLP model (e.g., spaCy or transformers) to identify intents in customer queries.

Example:
Write a function to simulate an AI agent processing a query.
Write a function to query a mock database and return an answer.

Deliverables:
Architecture diagram (brief description).
A code snippet (1-2 pages) that implements:
A simple query handling function.
An NLP-based intent recognition and mock database query.

```