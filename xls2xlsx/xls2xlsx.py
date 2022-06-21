# -*- coding: utf-8 -*-

import os
import win32com.client as win32
from openpyxl import load_workbook, Workbook


# 把一个文件夹内所有.xls文件改名为.xlsx, 并把原来xls文件删去
def f_xlsx(path):
    excel = win32.gencache.EnsureDispatch('Excel.Application')
    fld = os.listdir(path)
    for i in fld:
        if i[-4:] == ".xls":
            print("cover to:", i)
            i = path + i
            j = i + "x"
            wb = excel.Workbooks.Open(i)
            wb.SaveAs(j, FileFormat=51)
            os.remove(i)
            wb.Close()
    excel.Application.Quit()


# 把一个高级文件夹内所有xls转换为xlsx，包含子文件夹甚至孙文件夹中的文件，同时删除原来的xls文件
def f_xlsx_root(rootdir):  # rootdir:高级的含有内层文件夹的文件夹
    excel = win32.gencache.EnsureDispatch('Excel.Application')  # 启动excel程序

    for parent, dirnames, filenames in os.walk(rootdir):
        for fn in filenames:
            filedir = os.path.join(parent, fn)  # 每个文件的文件名，含路径：
            if filedir[-4:] == ".xls":
                print("cover to:", filedir)
                wb = excel.Workbooks.Open(filedir)  # 打开文件，带路径
                wb.SaveAs(filedir.replace('xls', "xlsx"), FileFormat=51)  # 文件另存为xlsx
                os.remove(filedir)  # 删除原来的xls文件
                wb.Close()
    excel.Application.Quit()  # 关闭excel程序


if __name__ == "__main__":
    f_xlsx_root("绝对路径")
