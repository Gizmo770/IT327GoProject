
/**  File: ParallelSum.java
 *  A team of threads adds up the numbers of an array
 * Name:
 * Date:
 *          Activity 13
 * IT 386
  */
import java.util.*;
import java.util.concurrent.locks.ReentrantLock;

public class ParallelSum
{
    // shared variables
    static int nThreads, N;
    static int arrayA[];
    static long totalSum = 0;

    public static void main(String[] args) throws InterruptedException
    {
        long startTime = System.currentTimeMillis();

        if (args.length > 0)
        {
            nThreads = Integer.parseInt(args[0]);
            N = Integer.parseInt(args[1]);
            System.out.print("Array size: " + N + ", ");
            System.out.println("Number of threads: " + nThreads);
        }
        else
        {
            System.out.println("Usage: java className <number of Threads> < N > ");
            System.exit(1);
        }

        arrayA = new int[N];
        for (int i = 0; i < N; i++)
        {
            arrayA[i] = i + 1;
        }
        /* Main thread prints initial array */
        if (N < 20)
        {
            System.out.println(Thread.currentThread().getName() + ": " + Arrays.toString(arrayA));
        }

        /* Create array to hold team of threads */
        Thread[] workers = new Thread[nThreads];
        int work = N / nThreads; // amount of work each thread will do
        int remainder = N % nThreads;
        for (int i = 0; i < nThreads; i++)
        {
            int low = i * work;
            int high = (i + 1) * work;
            // if statement
            if (remainder != 0 && i == nThreads - 1) {
                high += remainder;
            }
            Runnable obj = new Worker(low, high);
            workers[i] = new Thread(obj);
            workers[i].start();
        }

        for (int i = 0; i < nThreads; i++)
        {
            workers[i].join();
        }

        System.out.println("SeqSum " + seqSum() + ", Parallel Sum " + totalSum);

        long elapsedTime = System.currentTimeMillis() - startTime;
        System.out.println("Elapsed time: " + elapsedTime + " ms");
    } /* main */

    public static class Worker implements Runnable
    {
        static ReentrantLock mutex = new ReentrantLock();
        private int low, high;
        private long localSum;

        public Worker(int low, int high)
        {
            this.low = low;
            this.high = high;
        }

        public void run()
        {
            localSum = 0;
            for (int i = low; i < high; i++)
            {
                localSum += arrayA[i];
            }
            mutex.lock();
            try
            {
                totalSum += localSum;
            }
            finally
            {
                mutex.unlock();
            }

        }
    }

    public static long seqSum()
    {
        long sum = 0;
        for (int i = 0; i < N; i++)
        {
            sum += arrayA[i];
        }
        return sum;
    }

}
