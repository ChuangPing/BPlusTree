import pandas as pd
import matplotlib.pyplot as plt

# import numpy as np

binary_info = pd.read_csv(
    filepath_or_buffer="F:/BPlusTree/bptree/binarySearch.txt",  # 文件路径+名称
    encoding="ANSI",  # 编码方式
)

bpTree_info = pd.read_csv(
    filepath_or_buffer="F:/BPlusTree/bptree/BPlusTree.txt",  # 文件路径+名称
    encoding="ANSI",  # 编码方式
)
# print("info:\n", binary_info)
# print("info的类型：\n", type(binary_info))

# 设置列名
columns = ["key", "count"]
#  设置列名称
binary_info.columns = columns
bpTree_info.columns = columns
# print("keys:\n", info["count"].values)

# 读取数据
binary_count = binary_info["count"].values
search_key = binary_info["key"].values

bpTree_count = bpTree_info["count"].values
bpTree_key = bpTree_info["key"].values

# 创建DataFrame
bpTree_search_data = pd.DataFrame(
    data={
        "binaryCount": binary_count,
        "bpTreeCount": bpTree_count
    },
    index=search_key
)
# plt.hist(bpTree_search_data["bpTreeCount"])
# 二分查找
bpTree_search_data.plot(kind="bar", title="binarySearch vs BPlusTreeSearch")
# bpTree_search_data["bpTreeCount"].plot(kind="hist")
print("bpTree_search_data:\n", bpTree_search_data)

bpTree_than_binary = bpTree_search_data["bpTreeCount"] - bpTree_search_data["binaryCount"]
bpTree_than_binary_data = pd.DataFrame(
    data= {
        "bpTree_than_binary": bpTree_than_binary
    },
    index=search_key
)
bpTree_than_binary_data.plot(kind="bar", title="bpTree Than binary")

# 保存图片
# plt.savefig('F:/BPlusTree/bptree/image/B+树与二分查找对比图(500个测试用例).png')
# plt.savefig('F:/BPlusTree/bptree/image/B+树查找次数与二分查找次数之差（500个测试用例）.png')
plt.show()
