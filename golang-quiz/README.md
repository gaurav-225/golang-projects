# Quiz App in Go

Simple quiz application built in Go that asks users questions and evaluates their answers. 
It supports a timeout mechanism and reads questions from a CSV file.

# Run the Application

### With Custom Arguments
```bash
go run main.go -t 60 -f quiz.csv
```
- `-t 60`: Specifies the time limit for each question (in seconds).  
- `-f quiz.csv`: Specifies the CSV file containing the quiz questions.

### With Default Values
```bash
go run main.go
```
- Default time limit: 30 seconds.  
- Default file: `quiz.csv`.

---

# Learnings

1. **Using `time.NewTimer` for Timeout Handling**:  
   - used with `select` to detect timeout via the timer's `.C` channel.

2. **Handling User Responses in Goroutines**:  
   - User input is handled in a separate goroutine, which sends the responses to a channel.  
   - This allows concurrent processing of user input and timeout detection.

3. **Closing Channels After Processing**:  
   - After iterating through all the problems, the channel used for user responses is closed.

4. **Reading a CSV File in Go**:  
   - Open the file using `os.Open`.  
   - Pass the file object to `csv.NewReader`.  
   - Use `ReadAll()` method on reader to get the CSV file's content as a slice of records.

> I followed the YouTube tutorial: [Built A Golang Quiz App](https://www.youtube.com/watch?v=TVHGxz6tn2M)