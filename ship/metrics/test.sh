#!/bin/bash
#---
# Excerpted from "Effective Go Recipes",
# published by The Pragmatic Bookshelf.
# Copyrights apply to this code. It may not be used to create training material,
# courses, books, articles, and the like. Contact us if you are in doubt.
# We make no guarantees that this code is fit for any purpose.
# Visit https://pragprog.com/titles/mtgo for more book information.
#---

# OK call
curl -i 'http://localhost:8080/lookup?lat=1&lng=2'

# Error (invalid latitude)
curl -i 'http://localhost:8080/lookup?lat=100&lng=2'

# Check metrics
curl 'http://localhost:8080/debug/vars'