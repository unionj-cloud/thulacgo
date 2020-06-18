package thulacgo

import (
	"fmt"
	"strings"
)

func ExampleSegTest() {

	lac := NewThulacgo("models", "", false, false, false, byte('_'))
	defer lac.Deinit()
	ret := lac.Seg("新京报讯（记者 周萧）在北京时间今晨客场与赫塔菲的对决里，长时间人数处于劣势的“鹦鹉军团”0比0逼平对手。两轮比赛拿到4分，西班牙人终于摆脱了积分榜副班长的位置，反超莱加内斯1分排名倒数第二位。武磊赛后在个人社交媒体上用“所有人都是好样的”来评价西班牙人刚经历的比赛。")

	fmt.Println(strings.TrimSpace(ret))

	//Output:
	//新京报_nz 讯_g （_w 记者_n  _w 周萧_np ）_w 在_p 北京_ns 时间_n 今晨_t 客场_n 与_p 赫塔菲_np 的_u 对_p 决里_n ，_w 长_a 时间_n 人数_n 处于_v 劣势_n 的_u “_w 鹦鹉_n 军团_n ”_w 0_m 比_p 0_m 逼_v 平_a 对手_n 。_w 两_m 轮_q 比赛_v 拿_v 到_v 4分_t ，_w 西班牙人_n 终于_d 摆脱_v 了_u 积分榜_n 副_a 班长_n 的_u 位置_n ，_w 反_d 超_v 莱加内斯_ns 1分_t 排名_v 倒数_v 第二_m 位_q 。_w 武磊_np 赛后_t 在_p 个人_n 社交_n 媒体_n 上_f 用_p “_w 所有_a 人_n 都_d 是_v 好样_a 的_u ”_w 来_v 评价_v 西班牙_ns 人_n 刚_d 经历_v 的_u 比赛_v 。_w
}

func ExampleSegToSliceTest() {

	lac := NewThulacgo("models", "", false, false, false, byte('_'))
	defer lac.Deinit()
	ret := lac.SegToSlice("新京报讯（记者 周萧）在北京时间今晨客场与赫塔菲的对决里，长时间人数处于劣势的“鹦鹉军团”0比0逼平对手。两轮比赛拿到4分，西班牙人终于摆脱了积分榜副班长的位置，反超莱加内斯1分排名倒数第二位。武磊赛后在个人社交媒体上用“所有人都是好样的”来评价西班牙人刚经历的比赛。")

	result := strings.Join(ret, " ")
	fmt.Println(strings.TrimSpace(result))

	//Output:
	//新京报_nz 讯_g （_w 记者_n  _w 周萧_np ）_w 在_p 北京_ns 时间_n 今晨_t 客场_n 与_p 赫塔菲_np 的_u 对_p 决里_n ，_w 长_a 时间_n 人数_n 处于_v 劣势_n 的_u “_w 鹦鹉_n 军团_n ”_w 0_m 比_p 0_m 逼_v 平_a 对手_n 。_w 两_m 轮_q 比赛_v 拿_v 到_v 4分_t ，_w 西班牙人_n 终于_d 摆脱_v 了_u 积分榜_n 副_a 班长_n 的_u 位置_n ，_w 反_d 超_v 莱加内斯_ns 1分_t 排名_v 倒数_v 第二_m 位_q 。_w 武磊_np 赛后_t 在_p 个人_n 社交_n 媒体_n 上_f 用_p “_w 所有_a 人_n 都_d 是_v 好样_a 的_u ”_w 来_v 评价_v 西班牙_ns 人_n 刚_d 经历_v 的_u 比赛_v 。_w
}

func ExampleSegOnlyTest() {

	lac := NewThulacgo("models", "", true, false, false, byte('_'))
	defer lac.Deinit()
	ret := lac.Seg("尽管未能取得进球，但武磊的表现受到当地媒体肯定——“他在场上一次又一次地跑动，从不停下来。他现在与队友磨合得越来越好，不再拘谨了”。")

	fmt.Println(ret)

	//Output:
	//尽管 未能 取得 进球 ， 但 武磊 的 表现 受到 当地 媒体 肯定 —— “ 他 在 场上 一 次 又 一 次 地 跑动 ， 从不 停 下来 。 他 现在 与 队友 磨合 得 越来越 好 ， 不再 拘谨 了 ” 。
}

func ExampleUserpathTest() {

	lac := NewThulacgo("models", "userdict.txt", false, false, false, byte('_'))
	defer lac.Deinit()
	ret := lac.Seg("国家卫健委：16日31个省市和新疆生产建设兵团报告新增新冠肺炎确诊病例44例，境外输入病例11例，本土病例33例。")

	fmt.Println(ret)

	//Output:
	//国家_n 卫健委_j ：_w 16日_t 31_m 个_q 省市_n 和_c 新疆_ns 生产_v 建设_v 兵团_n 报告_v 新增_v 新冠肺炎_uw 确诊_v 病例_n 44_m 例_q ，_w 境外_s 输入_v 病例_n 11_m 例_q ，_w 本土_n 病例_n 33_m 例_n 。_w
}
