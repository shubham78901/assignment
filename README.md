
RUN :GO MOD TIDY
RUN:go build -o main ./api/cmd/main.go && ./main
 above command will generate binary and run

SWAGGER DOC AT :http://localhost:8000/swagger/index.html

<img width="1470" alt="Screenshot 2025-02-21 at 10 43 03 AM" src="https://github.com/user-attachments/assets/2967b69e-95f1-4f8b-8215-1e6d836e6d1d" />
