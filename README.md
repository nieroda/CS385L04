Replicas | Concurrency | Requests | Requests per second (mean) | Longest request | 98 percentile |
|--------|-------------|----------|----------------------------|-----------------|---------------|
|     1   | 50          | 10000    | 7234.92                    | 16              | 12            |
|    1    | 100         | 10000    | 7113.15                    | 27              | 21            |
|    1    | 500         | 10000    | 6020.22                    | 1446            | 1069          |
|    1    | 1000        | 10000    | 5328.10                    | 1856            | 1282          |
|    1    | 2000        | 10000    | 5161.17                    | 1929            | 1483          |
|    1    | 5000        | 20000    | 2408.22                    | 8294            | 4438          |
|    1    | 10000       | 30000    | 259.82                     | 115445          | 15086         |
|    1    | 15000       | 45000    | 373.00                     | 120561          | 116526        |

Replicas | Concurrency | Requests | Requests per second (mean) | Longest request | 98 percentile |
|--------|-------------|----------|----------------------------|-----------------|---------------|
|     3   | 50          | 10000    | 5852.59                    | 55              | 26            |
|    3    | 100         | 10000    | 5681.25                    | 99              | 66            |
|    3    | 500         | 10000    | 5542.43                    | 1256            | 280          |
|    3    | 1000        | 10000    | 5548.40                    | 1566            | 1147          |
|    3    | 2000        | 10000    | 4766.39                    | 2085            | 2008          |
|    3    | 5000        | 20000    | 4290.44                    | 4646            | 4439          |
|    3    | 10000       | 30000    | 3642.68                     | 8204          | 4813         |
|    3    | 15000       | 45000    | 628.98                     | 60848          | 60567        |




Replicas | Concurrency | Requests | Requests per second (mean) | Longest request | 98 percentile |
|--------|-------------|----------|----------------------------|-----------------|---------------|
|     6   | 50          | 10000    | 5978.73                   | 66              | 30            |
|    6    | 100         | 10000    | 6511.25                   | 151              | 80            |
|    6    | 500         | 10000    | 6536.03                   | 336            | 197          |
|    6    | 1000        | 10000    | 6694.09                    | 1255            | 1200          |
|    6    | 2000        | 10000    | 5719.95                   | 1581            | 1557          |
|    6    | 5000        | 20000    | 5992.19                    | 2353            | 1688          |
|    6    | 10000       | 30000    | 2143.52                     | 13678          | 11017         |
|    6    | 15000       | 45000    | 645.10                     | 60608          | 54191        |


Replicas | Concurrency | Requests | Requests per second (mean) | Longest request | 98 percentile |
|--------|-------------|----------|----------------------------|-----------------|---------------|
|     9   | 50          | 10000    | 8181.95                    | 85              | 45            |
|    9    | 100         | 10000    | 8968.08                    | 80              | 49            |
|    9    | 500         | 10000    | 9562.67                    | 220            | 150          |
|    9    | 1000        | 10000    | 9278.89                    | 536            | 378          |
|    9    | 2000        | 10000    | 9769.46                    | 951            | 321          |
|    9    | 5000        | 20000    | 7044.57                    | 2792            | 1660          |
|    9    | 10000       | 30000    | 6625.34                     | 4477          | 3500         |
|    9    | 15000       | 45000    | 588.54                     | 60652          | 54664        |




Replicas | Concurrency | Requests | Requests per second (mean) | Longest request | 98 percentile |
|--------|-------------|----------|----------------------------|-----------------|---------------|
|     12   | 50          | 10000    | 8693.28                    | 61              | 39            |
|    12    | 100         | 10000    | 9257.53                    | 90              | 59            |
|    12    | 500         | 10000    | 9496.95                    | 325            | 199          |
|    12    | 1000        | 10000    | 9775.25                    | 496            | 279          |
|    12    | 2000        | 10000    | 9377.63                    | 648            | 460          |
|    12    | 5000        | 20000    | 9757.54                    | 1918            | 1516          |
|    12    | 10000       | 30000    | 1666.29                     | 16004          | 15939         |
|    12    | 15000       | 45000    | 644.80                     | 48766          | 44002        |



Did you get a noticeable improvement between 9 and 12 replicas? Explain this behavior.

5-10% overall improvement, not the 25% I was expecting.

What is the maximum throughput your service can provide? What is the maximum concurrency that you can support, while maintaining maximum throughput? How many replicas (approximately) do you need to support that concurrency level and throughput?

While not really the max, the 'max' throughput of this server was 9757.54. This was measured with 12 replica sets, 5000 concurrent requests with 20,000 requests total.

Suppose we expect our service to have a maximum number of 2000 concurrent users. We have a Service Level Agreement with our customers in which we guarantee a maximum response time up to 1.5 seconds, with a mean response of less than 300 milliseconds, with requests 98% of the time less than 600 ms. Can we fulfill those requirements? If so, how many replicas will be needed?

Based off the data we will be able to fulfill these requirements with ~12 replicas.




