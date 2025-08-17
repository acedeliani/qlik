# QLIK Cloud Assignment

#### Architecture

##### Main receiver service
I would start with having a component which is dedicated to receiving and storing the data. As the vendor is very popular, it is of great importance not to lose any incoming orders. It can always be processed at a later stage. Receiving and storing the data in a database should be its priority and care must be taken to ensure that the rate of data can be handled by this component. Otherwise, additional components with load balancers may be used to manage the data flow. For improving access times, the data might be stored in a cache (e.g. Redis) as well for other components to access it.

##### Processing service (for higher performance)
A processing component may be introduced to offload processing for the main receiver service. It reads unprocessed data, converts it and returns the result either directly to another service or to a cache for later retrieval. It might be order creation, order retrieval or other operations that are considered too time consuming for the recevier.

##### Summary
I think a simple 2-service architecture will suffice given the information provided. Even if I believe that it is important to properly design an application from the start, designing for high performance should also be done gradually to prevent over-engineered solutions. In fact, I would start out with a single monolith service, as done in this application and take it from there.
