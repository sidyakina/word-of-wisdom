# word-of-wisdom

## Task
Design and implement “Word of Wisdom” tcp server.  
• TCP  server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.  
• The choice of the POW algorithm should be explained.  
• After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.  
• Docker file should be provided both for the server and for the client that solves the POW challenge

## Credentials
Quotes about space are taken from https://www.thefactsite.com/100-space-facts/

## Algorithm
Algorithm based on HashCash algorithm (https://en.wikipedia.org/wiki/Hashcash) but have several differences:

• SHA2 was used instead of SHA1. SHA1 is a cryptographically broken. Also, SHA2 take more time for generate hash and this adds more time difficulty for solve challenge.
• HashCash is using mail info for challenge, but we haven't such info for our client. Random string with symbols a-z, A-z, 0-9 is using instead.

Algorithm steps:

1. Server generate random string with N symbols
2. Client is finding random string with M symbols matches the next condition: SHA2(server_string+client_string) has K first zero bits.
3. Server receive client string and validate it

Current values: N=20, M=10, K=20

Value N affects number possible server strings. 
If hackers would to prepare hashes for all possible server strings they need prepare 64^20 hashes. 
One hash takes about 778ms in average (test/algorithms_test.go) and such preparation will take much time. 
If we need increase this time we can increase N. 
It won't take more time for one attempt (test/algorithms_test.go) but significantly increase number of variants.

## Time
14.00 ---> 14.30
15.50 ---> 16.20
18.05 ---> 20.25
20.50 ---> 21.30
13.40 ---> 14.27
14.40 ---> 16.10
18.36 ---> 20.50
13.00 ---> 15.10
19.00 ---> 21.14
21.45 ---> 22.35


