package com.oop;

class Hero {

    String name; // 英雄名字

    float hp; // 血量

    float armor; //护甲

    int moveSpeed; //移动速度

    // 坑队友
    void keng () {
    	System.out.println("肯呢个");
    }
}

class TestHero {
    public static void main(String[] args) {
        // 创建英雄盖伦并且初始化
        Hero garen = new Hero();
        garen.name = "盖伦";
        garen.hp = 1001.1f;
        garen.armor = 12.1f;
        garen.moveSpeed = 111;
        System.out.println(garen.name+garen.hp+garen.armor+garen.moveSpeed);

        // 创建英雄提莫并且初始化
        Hero teemo =  new Hero();
        teemo.name = "提莫";
        teemo.hp = 383f;
        teemo.armor = 14f;
        teemo.moveSpeed = 330;
    }

}