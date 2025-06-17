# Query Expansion Backend for Matching Queries to JSON and XML Files

Inspired by [Uber Query Understanding Engine](https://www.uber.com/blog/uber-eats-query-understanding/).

## Overview 
Inspired by Uber's utilization of Query Expansion (QE) to improve the relevance of search results, this project explores a backend system designed to match user queries to structured data sources such as JSON and XML files.

Query Expansion is a technique commonly used in information retrieval systems to enhance the original user query with additional, semantically similar terms. This helps improve recall and ensures more comprehensive results. A key component of QE is the ability to find similar queries in real time, which requires an efficient way to measure and retrieve nearest neighbors in a high-dimensional vector space.


## How It Works

Computing similarity across all entries in real time can be computationally expensive. To address this, we utilize approximate nearest neighbor search techniques. This can be done by preprocessing (offline) of all available documents. When a query is posted, it is tokenized and matched in a map to relevant restaurant objects (labeled documents in this project).