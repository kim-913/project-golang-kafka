# This is a sample Python script.

# Press ⌃R to execute it or replace it with your code.
# Press Double ⇧ to search everywhere for classes, files, tool windows, actions, and settings.

import mysql.connector
import random
import string

mydb = mysql.connector.connect(
    host='localhost',
    user='root',
    password='Kzr54321',
    port='3306',
    database='test_project'
)

mycursor = mydb.cursor()
sql = """INSERT INTO test (id, name, address, Continent) VALUES (%s, %s, %s, %s)"""

min_data = 1
max_data = 5000000
each = 5000000
ID = random.sample(range(min_data, max_data + 1), each)


continent = ["North America", "Asia", "South America", "Europe", "Africa", "Australia"]
for i in ID:
    res = []
    res.append(i)
    name = ''
    name += "".join(random.choice(string.ascii_lowercase) for j in range(random.randint(10, 15)))
    res.append(name)
    address = ''
    rand_num = random.randint(1, 10000)
    address += str(rand_num) + ' '
    rand_string1 = ''
    rand_string1 += "".join(random.choice(string.ascii_lowercase) for j in range(2, 7))
    address += rand_string1 + ' '
    rand_string2 = ''
    rand_string2 += "".join(random.choice(string.ascii_lowercase) for j in range(15 - len(address), 20 - len(address)))
    address += rand_string2
    res.append(address)
    user_continent = random.choice(continent)
    res.append(user_continent)
    # print(res)
    mycursor.execute(sql, (res[0], res[1], res[2], res[3]))
    mydb.commit()