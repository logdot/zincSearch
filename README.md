# ZincSearch
This is a project I created for a job application.
It is split into 2 excecutables and one small API.
The backbone of the project is the [ZincSearch Search Engine](https://zincsearch.com/).

This project used technologies that were completely new to me in both the frontend and the backend.

The tecnologies used for the frontend where:
  - Vue
  - Tailwind (css)

And for the backend:
  - Go
  - ZincSearch

## Enrin
Enrin (ENron INdexer) is a utility program to populate the ZincSearch backend with data from the Enron Email Database.
It recursively traverses the enron database and parses each of the mail files.
Originally Enrin was singlethreaded, taking up to 30+ minutes to index the entirety of the enron database.
Now it is multithreaded and can index the whole database in around 5 minutes, but with a hefty memory requirement.


## Envi
Envi (ENron VIsualizer) is a utility webserver to display the enron mail data from ZincSearch.

It is responsive to multiple screensizes, but IMO it would still not be a super enjoyable experience.
Altough the responsiveness does at least make it useable instead of outright unbearable.

<img src="https://user-images.githubusercontent.com/34782839/215234474-5840a88b-8c13-4661-9eb3-561ba6ffe29a.png" width="800"/>

![image](https://user-images.githubusercontent.com/34782839/215235333-3c2871ac-566f-4be9-acfc-83b76a2fe782.png)

## API
This project exposes a simple API.
The main way you interact with it is through an Authentication object created from a Authenticate function.
You give it the address of the ZincSearch server, an index, and a username and password.
Using the Authentication object you can perform various actions such as search queries or ingesting new data.

## Improvements
While at the current moment the project is serviceable, it could still do with some improvements.

- [ ] Pagination:
  Currently whenever you search for a term from within Envi it will query for all of the results at once.
  This makes it so if you search a common term (e.g. "enron") it can take a couple seconds for the data to be processed.
  If pagination where to be used this issue could be completely fixed.
  
- [ ] Deduplication:
  In the enron database an email can appear multiple times.
  This is because if person `A` sends an email to person `B` each of them will have one copy of the email in the enron database.
  Currently Enrin does not account for this and will bloat the database with duplicates.
  
- [x] Multithreaded Indexing:
  The indexing of the enron database is currently single threaded making it take a long time to process the 600'000 emails in the enron database.
  Due to Go's strengths it should be possible to easily multithread the indexing and make the process much faster.
  
- [x] Better Parsing:
  The parsing that is currently implemented is very fast, but frankly doesn't work the best.
  The most glaring issue is that currently the body of emails have erronious newlines at the beginning.

## Optimizations
The indexing of a single file is relatively fast, but a lot of time (80%+) is spent waiting to be given the file handle.
Due to this there isn't much that can be done to speed up the process, apart from multithreading the program.
With the indexing becoming multithreaded (one goroutine per file) and with a cap of 500 goroutines by default the indexing time is greatly shortened.
The indexing time goes from 30+ minutes (~333 files a second) to 5±1 minutes (~2000 files a second).