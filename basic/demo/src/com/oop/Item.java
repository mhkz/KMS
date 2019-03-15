package com.oop;

class Item {
    String name;
    int price;


    public static void main(String[] args) {
    	// 血瓶 
        Item booldVail = new Item();
        booldVail.name = "血瓶";
        booldVail.price = 50;
        
        // 草鞋
        Item strawSandals = new Item();
        strawSandals.name = "草鞋";
        strawSandals.price = 300;
       
        // 长剑
        Item sword = new Item();
        sword.name = "长剑";
        sword.price = 350;

        System.out.println(booldVail.name + booldVail.price);
        System.out.println(strawSandals.name + strawSandals.price);
        System.out.println(sword.name + sword.price);
    }
}