# FileUpload

Simple server for file upload only with Golang batteries

# Routes

- GET `/uploads/:filename`

View a file

- DELETE `/uploads/:filename`

Deletes the given file from UPLOADS_DIR

- PUT `/uploads/`
    - Body: FormData('image')
    
Insert the given file
