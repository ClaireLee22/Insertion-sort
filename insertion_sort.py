def insertionSort(array):
    for i in range(1, len(array)):
        j = i
        while j > 0 and array[j-1] > array[j]:
            swap(j-1, j, array) 
            j -= 1
    return array

def swap(i, j, array):
    array[i], array[j] = array[j], array[i]


if __name__ == "__main__":
    array = [8, 2, 5, 4, 1]
    print(insertionSort(array))

"""
output: [1, 2, 4, 5, 8]
"""