def fibo(num: int):
    nums: list[int] = [0, 1]

    for i in range(1, 100):
        nums.append(nums[i] + nums[i - 1])

    if num in nums:
        print(f"O número {num} pertence a sequência de Finonacci!")
    else:
        print(f"O número {num} não pertence a sequência de Fibonacci!")

fibo(12)
