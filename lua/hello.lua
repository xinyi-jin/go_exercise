#!/usr/local/bin/lua
require("lua/Fruit")
-- Hello World
print("Hello Lua!!!")

-- 局部变量
local n = 123
local str = "123"
print(n, type(n))
print(str, type(str))

-- string integer 运算
local new_str = n .. str
print(n + str)
print(new_str)
print(#new_str)

-- table
Site = { "apple", "orange" }
Site["key"] = 123
print(Site[1])
print(Site["key"])

-- 循环
n = 0
while n < 10 do
    if n == 3 then
        print("n==3")
    else
        print(n)
    end
    n = n + 1
    -- n = (n or 0) + 1
end
print("\n")
--[[ for i = 1, 10 do
    if i == 3 then
        print("i==3")
        goto continue
    end
    print(i)
    ::continue::
end ]]

for i = 1, 10 do
    repeat
        if i == 3 then
            print("i==", i)
        else
            print(i)
        end
    until true
end

if 0 then
    print("0 ", true)
end

if 1 then
    print("1 ", true)
else
    print("1 ", false)
end
local a = "goto"
local b = 123
if b == nil then
    print("b==nil")
else
    if a == nil then
        print(a == nil)
    else
        print(a ~= nil)
    end
    print("b!=nil")
end


-- function 函数
local function getString(num)
    return num .. ""
end

local pp = getString(4546)
print(pp, type(pp), #pp)

local function max(a, b)
    if a > b then
        return a
    else
        return b
    end
end

print("10 和 100, 更大值为：", max(10, 100))

local res1, res2, res3 = string.find("welcome ZhengZhou", "come")
print(res1)
print(res2)
print(res3)

-- 可变参数
local function avg(...)
    local total
    for _, value in ipairs({ select(4, ...) }) do
        -- total = total + value
        total = (total or 0) + value
    end

    local num = { ... }
    for i = 1, #num do
        print(string.format("i == %d", num[i]))
    end
    return total / select("#", ...)
end

print(avg(1, 4, 1, 20))




local tab3 = {
    nil, nil, nil, 1, 1, 1, nil, 1, 1, 1,
    1, 1, nil, 1, 1, 1, 1, 1, nil, 1,
    1, 1, 1, 1.2, "1", nil, nil, 1, nil, "1",
}

print("tab3的长度", #tab3)

-- print(table.getn(tab3))

-- 变相整除
print(5 % 2)
print((5 - 5 % 2) / 2)

-- 逻辑运算符
if 5 > 3 and 3 < 5 then
    print("and")
    if 5 > 10 or 5 < 10 then
        print("or")
        if not (5 > 10) then
            print("not")
        end
    end
end

if (true) then
    print("true")
else
    print("false")
end

-- 三元运算符
local sy = (1 > 5 and { "true" } or { "false" })[1]
local sy2 = 1 > 5 and "true" or "false"
print(string.format("sy:%q", sy))
print(string.format("sy2:%q", sy2))

-- 字符串
str = "你好哇，火箭龟"
print(string.len(str))
-- 报错 nil
-- print(utf8.len(str))

str = "ABCD under DD"
print(string.upper(str))
print(string.lower(str))
print(string.reverse(str))
local newstr, n = string.gsub(str, "d", "D", 1)
print(newstr, n)
local start, endd = string.find(str, "D", -1)
print(start, endd)

newstr = string.rep(str, 1)
newstr = "4564351312"
print(str, newstr)
newstr = string.sub(str, 4, 13)
print(newstr)

-- 数组
local arr = { 10, 20, 30, 40, 50 }
print("arr len:", #arr)
for i = 1, #arr do
    print(string.format("arr[%d] = %d", i, arr[i]))
end

arr[-1] = 1000
arr[0] = 100

for i = -2, 5, 1 do
    print(arr[i])
    -- print(string.format("arr2[%d] = %d", i, arr[i]))
end
-- 追加元素
arr[#arr + 1] = 60
print(arr[#arr])
table.remove(arr, 5)
for i = 1, #arr do
    print(string.format("arr3[%d] = %d", i, arr[i]))
end

-- table concat
local tmp = { "one", "two", "three" }
print(table.concat(tmp))
print(table.concat(tmp, ",", 1, #tmp))
print(table.concat(tmp, ",", 2, #tmp))

-- table insert remove
print(table.insert(tmp, "four"))
print(table.insert(tmp, 2, "bigbang"))
print(table.concat(tmp, ",", 1, #tmp))

print(table.remove(tmp, #tmp))
print(table.remove(tmp, 2))
print(table.concat(tmp, ",", 1, #tmp))

for k, v in ipairs(arr) do
    print("k:", k, "v:", v)
end
-- table sort
table.sort(arr, function(a, b)
    return a > b
end)
for k, v in ipairs(arr) do
    print("k:", k, "v:", v)
end

print("Fruit", Fruit.GetName())
print("Fruit local", Fruit.GetLocalName())

-- 垃圾回收
print(collectgarbage("count"))
tmp = nil
print(collectgarbage("count"))
collectgarbage("collect")
print(collectgarbage("count"))
