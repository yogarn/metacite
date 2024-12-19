# metacite
> A simple tool that takes a URL and generates an APA citation. No need to dig through metadata or format anything manually — just give it the URL, and you're good to go.

---

## Installation
```bash
# Step 1: Make sure you have Go installed
# Download Go here: https://golang.org/dl/

# Step 2: Clone the repository
git clone https://github.com/yogarn/metacite.git
cd metacite

# Step 3: Install dependencies
go mod tidy

# Step 4: Build the project
go build -o metacite .
```
## Usage
```
./metacite <url>
```
or run it directly without build
```
go run . <url>
```
