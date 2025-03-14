

---
# Note: Make sure **Go is installed** on your machine before running this command.  
# You can download Go from 👉 [golang.org/dl](https://go.dev/).  

```
# Setting Up GitHub Repository & Go Modules for the E-Commerce Project 🚀

This guide will walk you through:  
✅ Creating a GitHub repository  
✅ Cloning the repo locally  
✅ Setting up Go modules (`go mod init` & `go mod tidy`)  

---

## **1. Create a GitHub Repository** 🏗️  

1. Go to [GitHub](https://github.com/) and **log in**.  
2. Click on the **"New Repository"** button or visit:  
   👉 [Create a new repository](https://github.com/new)  
3. Fill in the details:  
   - **Repository Name:** `ecommerce-golang`  
   - **Description:** "E-Commerce project built with Golang, PostgreSQL, Redis, and Kafka"  
   - **Visibility:** Public or Private (your choice)  
   - **Initialize with README:** ✅ (Optional)  
4. Click **"Create Repository"** 🎉  

---

## **2. Clone the Repository Locally** 💻  
Now, copy the **repository URL** (from GitHub) and open your terminal.

```bash
git clone https://github.com/YOUR_USERNAME/ecommerce-golang.git
```
Replace `YOUR_USERNAME` with your actual GitHub username.  
This will download the project to your local machine.

Move into the project folder:
```bash
cd ecommerce-golang
```

---

## **3. Initialize Go Modules (`go mod init`)** 🛠️  

> **Note:** Make sure **Go is installed** on your machine before running this command.  
> You can download Go from 👉 [golang.org/dl](https://golang.org/dl/).  

To manage dependencies, initialize Go modules:  
```bash
go mod init github.com/YOUR_USERNAME/ecommerce-golang
```

Check if `go.mod` was created:  
```bash
ls
```
You should see a file named **`go.mod`** inside the project.
```

---

