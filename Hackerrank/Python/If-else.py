#!/bin/python3

import math
import os
import random
import re
import sys


num=int(input())
if 1==(num%2):
    print("Weird")
elif (num%2==0 and 1<=num<=5):
    print ("Not Weird")
elif (num%2==0 and 6<=num<=20):
    print ("Weird")
elif (num%2==0 and 20<num):
    print("Not Weird")
else :
    print ("-")