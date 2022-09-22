import pandas as pd
import matplotlib.pyplot as plt


def readFile(num):
    num = str(num)
    binary_info = pd.read_csv(
        filepath_or_buffer="F:/BPlusTree/bptree/data/binarySearch" + num + ".txt",  # 文件路径+名称
        encoding="ANSI",  # 编码方式
    )

    bpTree_info = pd.read_csv(
        filepath_or_buffer="F:/BPlusTree/bptree/data/BPlusTree" + num + ".txt",  # 文件路径+名称
        encoding="ANSI",  # 编码方式
    )
    return binary_info, bpTree_info

# 二分查找测试数据查找和
binary_sum_all = []
# b+树测试数据查找和
bpTree_sum_all = []
# 测试用例数量
test_sum = []
def creatDat(nums):
   for num in nums:
        binary_info, bpTree_info = readFile(num)
        # 设置列名
        columns = ["key", "count"]
        #  设置列名称
        binary_info.columns = columns
        bpTree_info.columns = columns
        # 将sum数量测试用例的所有比较次数相加
        binary_sum = binary_info["count"].sum()
        binary_sum = binary_sum / len(binary_info["count"])
        binary_sum_all.append(binary_sum)
        bpTree_sum = bpTree_info["count"].sum()
        bpTree_sum = bpTree_sum / len(bpTree_info["count"])
        bpTree_sum_all.append(bpTree_sum)
        test_sum.append(num)

testDataCount = [50, 100, 300, 500, 1000]
creatDat(testDataCount)
bpTree_search_sumData = pd.DataFrame(
    data={
        "binarySumAll": binary_sum_all,
        "bpTreeSumAll": bpTree_sum_all
    },
    index = test_sum
)
bpTree_search_sumData.plot(kind="bar", title="bpTree and binary average search count")
plt.savefig('F:/BPlusTree/bptree/image/B+树与二分查平均查找次数对比图(50, 100, 300, 500 测试用例).png')

# testDataCount = [50, 100, 300, 500, 1000, 5000, 10000]
# creatDat(testDataCount)
# bpTree_search_sumData = pd.DataFrame(
#     data={
#         "binarySumAll": binary_sum_all,
#         "bpTreeSumAll": bpTree_sum_all
#     },
#     index = test_sum
# )
# bpTree_search_sumData.plot(kind="bar", title="bpTree and binary average search count")
# plt.savefig('F:/BPlusTree/bptree/image/B+树与二分查平均查找次数对比图(50, 100, 300, 500, 1000, 5000, 10000 测试用例).png')

plt.show()

print("bpTree_search_sumData", bpTree_search_sumData)
# print(binary_info, bpTree_info)
