go func() {
sig := <-shutdownSignal
fmt.Printf("Получен сигнал завершения: %v\n", sig)

fmt.Println("Закрытие соединения с NATS Streaming...")
natsCon.Close()

fmt.Println("Закрытие соединения с базой данных...")
testDB.Close(ctx)

fmt.Println("Завершение программы")
os.Exit(0)
}()
