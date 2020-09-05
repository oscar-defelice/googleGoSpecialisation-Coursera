## Week 1 Quiz

1. If two tasks are executing in parallel, which of the following statements is true?

![](https://user-images.githubusercontent.com/49638680/92257220-271edc00-eed5-11ea-91fa-90f3cbafa646.png)

_Parallelism_ means two (or more) tasks can run at the very same time.
Since a processor only executes one instruction at time, you need different processors in order to execute more than one task in the same moment.

2. What does the von Neumann bottleneck state about computer architectures?

![](https://user-images.githubusercontent.com/49638680/92257224-28500900-eed5-11ea-9f53-786c3930b13c.png)

Roughly speaking, memory access is slower than the working rate of the CPU. In other words, we can talk about limited _throughput_ (data transfer rate) between the central processing unit (CPU) and memory compared to the amount of memory. Because the single bus can only access a limited portion of memory at a time, throughput is lower than the rate at which the CPU can work.
Hence, often CPU waits for the memory to "send" data.

3. What does Moore’s law directly observe?

![](https://user-images.githubusercontent.com/49638680/92257226-29813600-eed5-11ea-9c05-0fefc2758bae.png)

Moore's law is the observation that the number of transistors in a dense integrated circuit doubles about every two years.
This law is not a physical law, rather it is an observation and a projection of a historical trend.

4. How is dynamic power consumption related to voltage swing?

![](https://user-images.githubusercontent.com/49638680/92257228-2a19cc80-eed5-11ea-9347-e372e90ea62e.png)

As seen dynamic power is directly proportional to the square of the voltage swing, to the capacitance, to ratio of time switching and to clock frequency.

$$ P = \alpha C f V^2 $$

Hence, referring to voltage swing, the answer is the square of voltage variation.

5. Why can’t Dennard Scaling continue forever?

![](https://user-images.githubusercontent.com/49638680/92257231-2a19cc80-eed5-11ea-8da1-d4d01b804106.png)

Dennard observes that transistor dimensions could be scaled by -30% (0.7x) every technology generation, thus reducing their area by 50%. This would reduce circuit delays by 30% (0.7x) and therefore increase operating frequency by about 40% (1.4x). Finally, to keep the electric field constant, voltage is reduced by 30%, reducing energy by 65% and power (at 1.4x frequency) by 50%.

This cannot continue forever, of course, because if voltage swing goes under threshold voltage, physically the transistors will not switch.
Furthermore, the signal may become undistinguishable from noise.

6. What factor limits clock rates in future designs?

![](https://user-images.githubusercontent.com/49638680/92257234-2ab26300-eed5-11ea-978e-13394f762ccf.png)

We have seen how reducing dimensions finally means increasing operating temperature, and operating power. This limits the clock rates.
Furthermore, one has to consider that communications between different components of the CPU is not immediate because of the speed of light.

7. One benefit of concurrent execution on a single processor is that it can hide latency. What does this mean?

![](https://user-images.githubusercontent.com/49638680/92257237-2ab26300-eed5-11ea-9122-24d62dabd1ca.png)

Often (recall von Neumann bottleneck) a great amount of time is spent waiting for something (data, other process finishing, etc.).
Concurrency is an advantage because it allows to execute some task while another is waiting for something else.
