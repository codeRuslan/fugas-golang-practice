curl -X PUT \
http://localhost:8000/books \
-H 'Content-Type: application/json' \
-d '[
{
"name": "The Fellowship of the Ring",
"author": "J. R. R. Tolkien",
"year": 1954
},
{
"name": "The Two Towers",
"author": "J. R. R. Tolkien",
"year": 1954
},
{
"name": "The Return of the King",
"author": "J. R. R. Tolkien",
"year": 1955
}
]'
