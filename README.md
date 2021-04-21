# Kanal
## Overview

## Endpoints
- [ ] List news
  - [ ] Filter by status
  - [ ] Filter by topic
- [ ] Add news
- [ ] Update news
- [ ] Remove news
- [ ] List tags
- [ ] Add tags
- [ ] Update tags
- [ ] Remove tags

## Models
### Topic
- Several news belongs to a topic
- Has an id and a name

### News
- Has an id, title, author, editor, published date, and article body
- Each news belongs only to one topic
- Each news can have one or more Tags

### Tags
- Has an id and a name
- Many tags belongs to many news

### ER Diagram
[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgVE9QSUMgfHwuLm97IE5FV1MgOiBcImJlbG9uZ3MgdG9cIlxuICAgIE5FV1MgfXwgLi4gb3sgVEFHUyA6IFwiaGFzXCJcblxuICAgICAgICAgICAgIiwibWVybWFpZCI6e30sInVwZGF0ZUVkaXRvciI6ZmFsc2V9)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgVE9QSUMgfHwuLm97IE5FV1MgOiBcImJlbG9uZ3MgdG9cIlxuICAgIE5FV1MgfXwgLi4gb3sgVEFHUyA6IFwiaGFzXCJcblxuICAgICAgICAgICAgIiwibWVybWFpZCI6e30sInVwZGF0ZUVkaXRvciI6ZmFsc2V9)

```
erDiagram
    TOPIC ||..o{ NEWS : "belongs to"
    NEWS }| .. o{ TAGS : "has"
```

## Assumptions
- Assuming each news belongs to only one topic
- Assuming we're building for a news company, so we don't have to record `publisher` information on News model.
- Assuming we don't have to worry about multimedia objects/blobs (e.g. videos, images, etc)
- Assuming we don't have to worry about links on the News's article

## TODO
- [ ] DB Models
  - [ ] Topic
  - [ ] News
  - [ ] Tags
- [ ] Endpoints
- [ ] Tests