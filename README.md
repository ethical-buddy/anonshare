ANONSHARE

backend--test

go run main.go-- cd into backend
 for upload api , curl request is 


 curl -X POST http://localhost:5000/upload \
  -H "Content-Type: application/json" \
  -d '{
    "hash": "testhash1234566464657",
    "size": "1.2 MB",
    "time": "2025-07-10T12:30:00Z",
    "type": "text/plain",
    "peers": [
      {
        "peer_id": "peer123",
        "ip": "192.168.1.100",
        "port": "8080",
        "file_path": "/home/user/file.txt",
        "filename": "file.txt",
        "description": "Test file"
      }
    ]
  }' 

  for files api --  curl -X GET "http://localhost:5000/files" 

for download api -----
                    curl -X GET http://localhost:5000/download \
  -H "Content-Type: application/json" \
  -d '{"hash": "testhash123"}'   --- first upload to see the results









