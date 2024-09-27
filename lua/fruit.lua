#!/usr/local/bin/lua

-- 定义模块名称
Fruit = {
    name = "orange",
    price = 0
}

-- 定义模块属性
MODEL_NAME = "MODEL_NAME"

local name = "local fruit"
local function getName()
    return Fruit.name
end

function Fruit.GetName()
    return getName()
end

function Fruit.GetLocalName()
    return MODEL_NAME
end

return Fruit
