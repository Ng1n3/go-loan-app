# 💸 Go Loan Bank

A simple command-line banking application built in Go that handles basic loan management and account operations

## 🎨 Features
- Create and manage bank accounts
- Check loan status and eligibility
- Process loan applications up to 5000 token
- Track and process loan repayments
- Account management and deletion
- Account number lookup by first name
- Phone number validation

## 📝 Prerequisites
- Go 1.19 or higher
- Make (for using the Makefile)

## 🚀 Installation
1. Clone the respository
   ```bash
    git clone https://github.com/Ng1n3/go-loan-app
    cd go-loan-bank
   ```
2. Install dependencies
   ```bash
    go mod download
   ```

## 🧱 Project Structure
```
├── account/
│   └── account.go       # Account structure and methods
├── fileops/
│   └── fileops.go       # File operations for data persistence
├── menu/
│   └── menu.go         # CLI menu and user interaction
├── main.go             # Application entry point
└── Makefile           # Build and run automation
```
## 💪 Usage
### Running the Application
Use the Makefile to build and run the application:
```bash
make run
```
This will:

- Format the code
- Run go vet
- Build the binary
- Execute the program

### Menu Options
The application presents a CLI menu with the following options:

1. Create an Account
2. Check Loan Status
3. Get a Loan
4. Repay Loan
5. Close an Account
6. Check account number
7. Exit

### Account Creation
When creating an account, you'll need to provide:

- First name
- Last name
- Phone number (must be exactly 10 digits)

The system will automatically:

- Generate a unique 4-digit account number
- Set initial loan amount available to 5000
- Set loan status to false
- Initialize loan counter and current loan amount

### Loan Management

- Maximum initial loan amount: 5000 tokens
- Loans are tracked per account
- Multiple loans are not allowed (must repay current loan first)
- Partial repayments are supported

### Data Persistence
Account data is stored in ``account.json`` file, which maintains:

- Personal information (name, phone number)
- Account details (account number, creation date)
- Loan information (status, available amount, current loan)

### Error Handling
The application handles various error cases:

- Invalid phone numbers
- Empty required fields
- Non-existent accounts
- Invalid loan amounts
- Outstanding loan checks

## Dependencies

- github.com/Pallinder/go-randomdata: For generating random account numbers

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (``git checkout -b feature/amazing-feature``)
3. Commit your changes (``git commit -m 'Add some amazing feature'``)
4. Push to the branch (``git push origin feature/amazing-feature``)
5. Open a Pull Request

## 💬 Contact
- 📨[Olayinka Emmanuel](code.with.muyiwa@gmail.com])
- [Github](https://github.com/Ng1n3)
- [X(Formerly known as Twitter)](https://x.com/34z1)

## 👏 Acknowledgments

- Thanks to the Go community for the excellent standard library
- The randomdata package for account number generation