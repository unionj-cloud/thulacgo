//
// Created by tqccc on 2020/6/24.
//

#include <cstdlib>
#include <cstdio>
#include "../thulac.h"

void ExampleSegTest() {
    Thulac lac = NewThulac("../models", "", 0, 0, 0, '_');
    const char *result = Seg(lac,
                             "新京报讯（记者 周萧）在北京时间今晨客场与赫塔菲的对决里，长时间人数处于劣势的“鹦鹉军团”0比0逼平对手。两轮比赛拿到4分，西班牙人终于摆脱了积分榜副班长的位置，反超莱加内斯1分排名倒数第二位。武磊赛后在个人社交媒体上用“所有人都是好样的”来评价西班牙人刚经历的比赛。");

    printf("分词结果是：\n");
    printf(result);
    // 释放
    Deinit(lac);
}

int main() {

    system("chcp 65001");

    ExampleSegTest();

    return 0;
}

