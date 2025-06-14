package main

import (
	"fmt"
	"math"
)

/*
Example 1:
Input: tires = [[2,3],[3,4]], changeTime = 5, numLaps = 4
Output: 21
Explanation:
Lap 1: Start with tire 0 and finish the lap in 2 seconds.
Lap 2: Continue with tire 0 and finish the lap in 2 * 3 = 6 seconds.
Lap 3: Change tires to a new tire 0 for 5 seconds and then finish the lap in another 2 seconds.
Lap 4: Continue with tire 0 and finish the lap in 2 * 3 = 6 seconds.
Total time = 2 + 6 + 5 + 2 + 6 = 21 seconds.
The minimum time to complete the race is 21 seconds.

Example 2:
Input: tires = [[1,10],[2,2],[3,4]], changeTime = 6, numLaps = 5
Output: 25
Explanation:
Lap 1: Start with tire 1 and finish the lap in 2 seconds.
Lap 2: Continue with tire 1 and finish the lap in 2 * 2 = 4 seconds.
Lap 3: Change tires to a new tire 1 for 6 seconds and then finish the lap in another 2 seconds.
Lap 4: Continue with tire 1 and finish the lap in 2 * 2 = 4 seconds.
Lap 5: Change tires to tire 0 for 6 seconds then finish the lap in another 1 second.
Total time = 2 + 4 + 6 + 2 + 4 + 6 + 1 = 25 seconds.
The minimum time to complete the race is 25 seconds.
*/

func main() {
	// f, err := os.Create("profile.pprof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = pprof.StartCPUProfile(f)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer pprof.StopCPUProfile()

	// minimumFinishTime([][]int{{1, 2}}, 5, 4) // 21 seconds
	// minimumFinishTimeBottomUp([][]int{{2, 3}, {3, 4}}, 5, 4) // 21 seconds
	// minimumFinishTimeBottomUp([][]int{{1, 10}, {2, 2}, {3, 4}}, 6, 5) // 25 seconds
	minimumFinishTimeBottomUp([][]int{{96, 3}, {68, 2}, {53, 4}, {60, 8}, {29, 8}, {96, 8}, {31, 10}, {5, 4}, {49, 6}, {54, 7}, {90, 7}, {7, 7}, {97, 2}, {50, 9}, {34, 2}, {89, 7}, {51, 7}, {73, 3}, {42, 4}, {24, 7}, {99, 3}, {34, 10}, {33, 9}, {45, 7}, {32, 2}, {59, 2}, {76, 3}, {10, 6}, {78, 7}, {19, 4}, {65, 2}, {30, 9}, {10, 5}, {84, 5}, {62, 4}, {87, 2}, {59, 8}, {29, 5}, {40, 4}, {76, 6}}, 15, 71) // 1405
	// minimumFinishTimeBottomUp([][]int{{3, 4}, {84, 2}, {63, 8}, {72, 8}, {82, 7}, {83, 6}, {23, 2}, {77, 5}, {51, 10}, {28, 2}, {47, 9}, {8, 3}, {48, 3}, {56, 3}, {8, 10}, {66, 6}, {92, 9}, {44, 6}, {23, 5}, {5, 6}, {86, 9}, {13, 10}, {91, 3}, {2, 2}, {8, 4}, {67, 8}, {63, 6}, {52, 5}, {42, 10}, {3, 9}, {66, 5}, {35, 10}, {63, 6}, {65, 6}, {22, 8}, {40, 9}, {43, 4}, {73, 9}, {81, 5}, {32, 2}, {30, 5}, {80, 9}, {50, 4}, {35, 4}, {52, 7}, {11, 5}, {7, 8}, {68, 3}, {54, 8}, {49, 8}}, 90, 87) // 2526
	// 58999
	// minimumFinishTimeBottomUp([][]int{{757, 17}, {157, 7}, {592, 23}, {421, 9}, {956, 8}, {625, 30}, {830, 29}, {805, 23}, {638, 24}, {113, 4}, {145, 4}, {702, 25}, {334, 19}, {424, 9}, {185, 14}, {523, 24}, {601, 18}, {958, 7}, {705, 6}, {844, 24}, {699, 6}, {743, 4}, {806, 18}, {945, 26}, {112, 5}, {165, 16}, {228, 28}, {540, 15}, {356, 14}, {258, 12}, {111, 21}, {776, 20}, {833, 20}, {466, 18}, {829, 6}, {898, 12}, {260, 27}, {219, 11}, {356, 3}, {886, 24}, {714, 6}, {422, 18}, {784, 29}, {772, 10}, {659, 20}, {721, 7}, {105, 24}, {328, 16}, {266, 30}, {504, 17}, {920, 3}, {781, 11}, {40, 28}, {912, 15}, {361, 9}, {159, 22}, {594, 15}, {559, 6}, {911, 8}, {790, 22}, {813, 30}, {4, 2}, {884, 14}, {418, 28}, {895, 25}, {37, 2}, {327, 2}, {477, 8}, {36, 11}, {513, 3}, {624, 28}, {584, 24}, {417, 25}, {242, 22}, {677, 11}, {821, 12}, {674, 6}, {829, 25}, {597, 18}, {190, 25}, {348, 25}, {339, 23}, {736, 27}, {415, 21}, {745, 13}, {387, 5}, {847, 15}, {766, 9}, {893, 30}, {795, 23}, {471, 15}, {340, 15}, {700, 12}, {589, 10}, {765, 4}, {12, 5}, {439, 6}, {66, 22}, {620, 25}, {241, 19}, {526, 21}, {544, 7}, {509, 13}, {964, 19}, {110, 2}, {421, 6}, {449, 4}, {883, 30}, {733, 16}, {818, 2}, {894, 19}, {365, 12}, {63, 13}, {365, 25}, {735, 25}, {491, 10}, {794, 4}, {744, 4}, {846, 14}, {93, 25}, {896, 13}, {929, 11}, {565, 22}, {323, 24}, {623, 2}, {706, 5}, {457, 12}, {505, 6}, {175, 28}, {702, 6}, {902, 29}, {293, 24}, {976, 4}, {40, 16}, {646, 28}, {162, 14}, {197, 7}, {561, 9}, {345, 10}, {222, 8}, {428, 28}, {453, 9}, {318, 21}, {523, 21}, {311, 19}, {704, 11}, {679, 10}, {136, 17}, {883, 24}, {668, 15}, {132, 21}, {352, 20}, {592, 7}, {609, 23}, {656, 13}, {919, 10}, {751, 3}, {88, 19}, {986, 7}, {114, 4}, {652, 25}, {703, 22}, {381, 25}, {285, 10}, {478, 26}, {714, 7}, {206, 3}, {516, 18}, {705, 21}, {179, 29}, {846, 16}, {773, 6}, {62, 7}, {361, 23}, {587, 23}, {636, 23}, {186, 2}, {256, 17}, {609, 19}, {642, 30}, {432, 7}, {273, 10}, {297, 2}, {881, 12}, {74, 21}, {724, 10}, {922, 29}, {429, 27}, {377, 12}, {177, 2}, {550, 29}, {311, 27}, {46, 17}, {522, 5}, {876, 11}, {494, 12}, {181, 6}, {160, 16}, {322, 18}, {645, 6}, {384, 3}, {668, 27}, {172, 5}, {799, 14}, {677, 16}, {600, 5}, {875, 26}, {580, 26}, {904, 12}, {730, 9}, {796, 30}, {206, 20}, {889, 18}, {753, 15}, {925, 5}, {463, 28}, {224, 12}, {447, 2}, {964, 27}, {741, 3}, {222, 13}, {435, 24}, {344, 23}, {298, 7}, {876, 18}, {879, 10}, {18, 24}, {311, 13}, {620, 27}, {576, 16}, {631, 5}, {191, 9}, {181, 8}, {322, 6}, {973, 3}, {866, 29}, {763, 5}, {702, 22}, {303, 8}, {191, 3}, {273, 12}, {264, 2}, {752, 15}, {315, 22}, {299, 8}, {784, 19}, {526, 28}, {322, 28}, {670, 22}, {741, 24}, {374, 12}, {687, 17}, {960, 21}, {429, 7}, {429, 11}, {230, 11}, {714, 24}, {907, 21}, {283, 24}, {908, 6}, {104, 20}, {502, 2}, {899, 28}, {967, 5}, {15, 27}, {448, 30}, {873, 11}, {927, 7}, {647, 20}, {454, 27}, {14, 20}, {366, 20}, {764, 10}, {648, 29}, {307, 29}, {368, 17}, {156, 26}, {915, 20}, {967, 10}, {113, 24}, {457, 6}, {142, 12}, {989, 18}, {556, 12}, {243, 12}, {998, 23}, {442, 23}, {677, 27}, {686, 25}, {189, 3}, {208, 15}, {438, 11}, {254, 2}, {22, 29}, {164, 26}, {384, 18}, {732, 12}, {128, 10}, {83, 17}, {980, 2}, {997, 30}, {667, 25}, {704, 30}, {837, 4}, {537, 26}, {345, 10}, {538, 16}, {684, 4}, {763, 17}, {280, 17}, {739, 29}, {819, 12}, {137, 11}, {223, 6}, {866, 5}, {941, 2}, {438, 10}, {486, 30}, {972, 10}, {558, 5}, {748, 19}, {662, 19}, {814, 3}, {748, 8}, {944, 2}, {877, 19}, {129, 27}, {938, 9}, {522, 13}, {241, 14}, {328, 4}, {368, 7}, {890, 30}, {532, 26}, {810, 3}, {239, 19}, {260, 18}, {763, 21}, {637, 21}, {805, 21}, {351, 17}, {779, 24}, {569, 9}, {96, 9}, {706, 13}, {721, 9}, {172, 6}, {995, 15}, {104, 4}, {665, 10}, {80, 11}, {326, 8}, {11, 20}, {482, 17}, {948, 30}, {938, 14}, {767, 26}, {872, 16}, {522, 14}, {718, 30}, {671, 26}, {245, 24}, {92, 15}, {91, 13}, {410, 21}, {886, 23}, {577, 4}, {740, 18}, {595, 30}, {666, 4}, {30, 29}, {350, 7}, {333, 24}, {71, 7}, {302, 23}, {526, 14}, {720, 5}, {748, 15}, {186, 24}, {934, 24}, {332, 17}, {525, 12}, {196, 14}, {26, 23}, {655, 29}, {14, 21}, {481, 18}, {292, 10}, {67, 13}, {897, 17}, {243, 19}, {930, 18}, {770, 4}, {834, 22}, {983, 25}, {214, 27}, {783, 14}, {887, 26}, {566, 11}, {4, 16}}, 347, 630)
	// minimumFinishTime([][]int{{9869, 239}, {87121, 40}, {90665, 703}, {48770, 13}, {29059, 750}, {32932, 18}, {79723, 864}, {90389, 161}, {51421, 947}, {84550, 861}, {98102, 64}, {76961, 990}, {8584, 257}, {61605, 233}, {18994, 78}, {4040, 635}, {39032, 312}, {58120, 602}, {71247, 641}, {92547, 749}, {52193, 808}, {94641, 528}, {67703, 992}, {33182, 78}, {42220, 215}, {66598, 999}, {2898, 180}, {77820, 513}, {90071, 547}, {94742, 778}, {51107, 283}, {35797, 571}, {3621, 502}, {80918, 780}, {11618, 132}, {71859, 543}, {40984, 469}, {9330, 476}, {70253, 623}, {37756, 893}, {67500, 592}, {13419, 310}, {83245, 193}, {77138, 549}, {81473, 522}, {17453, 896}, {4001, 717}, {5004, 11}, {8361, 796}, {2146, 393}, {41753, 266}, {87883, 292}, {4656, 853}, {28954, 749}, {83260, 468}, {97671, 958}, {45480, 440}, {38771, 993}, {40570, 288}, {22455, 926}, {13275, 377}, {48029, 63}, {1823, 327}, {81201, 502}, {31, 119}, {6705, 570}, {38124, 477}, {33220, 776}, {59087, 129}, {10779, 861}, {83927, 721}, {50209, 945}, {16994, 393}, {64972, 178}, {95790, 852}, {87836, 762}, {31891, 919}, {89961, 345}, {97621, 354}, {15143, 13}, {2173, 494}, {56418, 205}, {15933, 904}, {84629, 363}, {23849, 658}, {24270, 923}, {14266, 814}, {68991, 167}, {69974, 721}, {48793, 849}, {73860, 691}, {25485, 933}, {81731, 684}, {1933, 382}, {74806, 325}, {16431, 554}, {90071, 426}, {89648, 476}, {57673, 437}, {42700, 97}, {39198, 21}, {35228, 258}, {74578, 720}, {77545, 469}, {12576, 502}, {37680, 91}, {82021, 383}, {36395, 872}, {38406, 673}, {90719, 290}, {47186, 138}, {58552, 413}, {94423, 115}, {13207, 452}, {17625, 189}, {49118, 283}, {48092, 693}, {94438, 499}, {21095, 40}, {99463, 86}, {17302, 58}, {81360, 25}, {90147, 32}, {15841, 525}, {39945, 588}, {89734, 578}, {52205, 303}, {86286, 968}, {66305, 666}, {73421, 50}, {41909, 381}, {3603, 558}, {40432, 945}, {2129, 162}, {32923, 585}, {15673, 386}, {73620, 361}, {1400, 788}, {52629, 876}, {89000, 300}, {97171, 331}, {49911, 442}, {70859, 801}, {63886, 354}, {98997, 970}, {94408, 219}, {13162, 28}, {62146, 19}, {18555, 685}, {53922, 718}, {68426, 821}, {22702, 399}, {13659, 963}, {7504, 884}, {69877, 589}, {75468, 968}, {84399, 20}, {44602, 214}, {56113, 905}, {21804, 818}, {91827, 78}, {88795, 576}, {68141, 897}, {70873, 713}, {73525, 404}, {80353, 360}, {3558, 671}, {17846, 548}, {69396, 898}, {40957, 55}, {68987, 78}, {51190, 678}, {45997, 199}, {44705, 91}, {5916, 389}, {20082, 28}, {63036, 945}, {53095, 207}, {28735, 453}, {73510, 792}, {14942, 20}, {6589, 86}, {68166, 936}, {30426, 4}, {71677, 237}, {95927, 530}, {62750, 914}, {14790, 336}, {32686, 241}, {30591, 573}, {16644, 80}, {21343, 331}, {13465, 346}, {35911, 268}, {29420, 687}, {46817, 810}, {35614, 728}, {18948, 203}, {7484, 525}, {9500, 123}, {37757, 323}, {29774, 49}, {11107, 225}, {49440, 940}, {6607, 940}, {9905, 899}, {91785, 251}, {59274, 842}, {38799, 937}, {22854, 723}, {98312, 503}, {86833, 93}, {12387, 338}, {14404, 365}, {39374, 772}, {16810, 244}, {30253, 565}, {39070, 530}, {36629, 871}, {60627, 509}, {15418, 463}, {82561, 562}, {30293, 681}, {20577, 784}, {53002, 596}, {74349, 750}, {32634, 579}, {28042, 783}, {1776, 773}, {21009, 490}, {21233, 146}, {67617, 210}, {22406, 839}, {82722, 3}, {52919, 136}, {59989, 431}, {10252, 434}, {80492, 618}, {20938, 307}, {24749, 453}, {22581, 737}, {50566, 333}, {91968, 850}, {36996, 435}, {19246, 716}, {90050, 644}, {25581, 240}, {27559, 782}, {7029, 567}, {73646, 83}, {77301, 534}, {37279, 465}, {17734, 418}, {95891, 577}, {69601, 889}, {87885, 302}, {36717, 871}, {92971, 199}, {95258, 557}, {15355, 539}, {83992, 123}, {72590, 29}, {62747, 524}, {52979, 883}, {17382, 759}, {73052, 749}, {14467, 497}, {5653, 27}, {38213, 413}, {64393, 194}, {18188, 510}, {47954, 625}, {60186, 290}, {70329, 808}, {96628, 959}, {23902, 20}, {66594, 715}, {98608, 314}, {69032, 466}, {11445, 488}, {55361, 432}, {44, 797}, {31797, 581}, {55231, 722}, {8660, 64}, {50434, 589}, {71064, 870}, {95810, 8}, {23604, 126}, {96759, 619}, {99158, 815}, {96703, 841}, {6470, 234}, {86288, 806}, {53390, 606}, {20789, 351}, {28528, 858}, {27982, 785}, {76194, 868}, {31623, 894}, {10603, 739}, {8682, 319}, {94661, 869}, {70747, 817}, {83960, 135}, {78441, 410}, {12861, 771}, {88197, 456}, {16700, 727}, {25051, 879}, {84112, 452}, {13438, 962}, {12693, 598}, {71874, 322}, {78110, 449}, {77457, 48}, {50058, 661}, {77833, 200}, {5989, 356}, {63168, 482}, {99037, 184}, {91209, 800}, {92252, 988}, {79181, 526}, {94958, 79}, {40545, 510}, {24105, 469}, {85220, 24}, {25440, 308}, {14741, 956}, {11047, 324}, {68642, 500}, {18404, 24}, {83254, 744}, {28690, 981}, {81360, 302}, {70465, 814}, {26797, 727}, {39761, 884}, {84519, 938}, {81753, 544}, {59875, 29}, {17158, 92}, {89362, 313}, {75937, 821}, {17987, 504}, {91083, 8}, {45155, 780}, {99044, 735}, {7489, 940}, {10859, 713}, {68374, 986}, {21913, 981}, {70719, 888}, {18860, 821}, {88970, 122}, {74241, 84}, {87405, 405}, {44734, 183}, {36711, 391}, {24293, 591}, {84266, 456}, {93939, 340}, {44972, 722}, {75633, 440}, {60238, 236}, {64738, 615}, {13582, 497}, {34880, 656}, {91011, 536}, {32157, 840}, {42356, 73}, {95943, 934}, {93982, 734}, {27526, 1000}, {17028, 971}, {10883, 582}, {63921, 668}, {16071, 680}, {60476, 756}, {52213, 238}, {48298, 183}, {72670, 51}, {22622, 664}, {62230, 46}, {89423, 772}, {40765, 158}, {56640, 748}, {61622, 127}, {44384, 998}, {68706, 684}, {42037, 339}, {15690, 149}, {45206, 988}, {98761, 560}, {37705, 56}, {42773, 465}, {50661, 177}, {74142, 755}, {80964, 505}, {82746, 515}, {10685, 676}, {5323, 353}, {71226, 772}, {4511, 448}, {9921, 719}, {52005, 850}, {7221, 57}, {87452, 33}, {8550, 512}, {69748, 486}, {12002, 1000}, {11344, 533}, {22517, 842}, {82425, 485}, {56855, 911}, {92354, 173}, {36453, 660}, {41582, 792}, {48439, 692}, {65174, 370}, {68204, 334}, {41391, 671}, {27447, 977}, {87548, 954}, {83774, 172}, {33527, 398}, {29436, 152}, {44062, 497}, {96131, 561}, {84584, 343}, {94871, 694}, {57733, 948}, {10044, 918}, {22318, 512}, {78341, 550}, {59346, 891}, {20309, 424}, {98700, 75}, {59349, 947}, {49102, 17}, {79829, 279}, {79728, 453}, {30860, 823}, {7457, 466}, {62346, 468}, {85282, 729}, {42062, 28}, {90265, 262}, {48508, 874}, {55440, 690}, {67224, 898}, {33863, 922}, {90385, 23}, {85358, 216}, {51250, 678}, {33463, 343}, {7058, 948}, {61735, 979}, {46933, 94}, {28528, 464}, {58243, 207}, {41291, 246}, {7757, 734}, {84123, 962}, {47709, 593}, {42090, 918}, {68484, 753}, {78477, 402}, {85158, 464}, {13696, 632}, {40683, 827}, {56492, 372}, {39380, 559}, {59971, 404}, {55568, 967}, {28445, 491}, {36274, 260}, {37016, 505}, {90931, 708}, {73361, 881}, {51812, 636}, {57683, 914}, {45313, 885}, {42923, 800}, {67774, 771}, {13178, 365}, {13679, 478}, {90459, 363}, {6018, 21}, {94647, 608}, {56729, 502}, {77874, 736}, {94805, 553}, {8061, 309}, {42845, 275}, {39228, 410}, {77842, 597}, {56010, 194}, {32928, 292}, {48016, 552}, {495, 818}, {20847, 801}, {67934, 867}, {76691, 113}, {75292, 535}, {48217, 820}, {83383, 649}, {48218, 261}, {19315, 701}, {98916, 546}, {85921, 490}, {51069, 134}, {2761, 893}, {82045, 101}, {76809, 411}, {28781, 891}, {74069, 65}, {95773, 179}, {8012, 32}, {38393, 781}, {580, 96}, {47645, 588}, {80886, 143}, {67798, 741}, {70117, 143}, {22671, 757}, {27306, 229}, {17321, 367}, {85802, 16}, {55463, 210}, {50678, 988}, {91000, 85}, {37477, 225}, {56064, 207}, {38201, 771}, {13550, 992}, {42178, 53}, {78524, 852}, {2452, 381}, {97147, 716}, {6995, 423}, {90314, 889}, {43279, 915}, {18422, 5}, {45386, 949}, {4424, 16}, {5518, 926}, {75532, 545}, {32810, 984}, {97149, 98}, {41844, 344}, {55724, 365}, {46381, 519}, {59398, 116}, {16704, 805}, {27294, 148}, {58039, 298}, {83268, 534}, {2025, 926}, {86362, 261}, {25866, 953}, {43126, 154}, {20000, 809}, {13926, 437}, {11307, 10}, {61846, 423}, {96542, 848}, {36487, 863}, {10196, 325}, {64570, 448}, {14973, 756}, {88327, 35}, {88157, 984}, {79131, 502}, {95687, 464}, {88252, 52}, {63454, 773}, {35119, 67}, {16606, 406}, {24715, 501}, {81593, 984}, {38310, 777}, {10330, 543}, {86306, 46}, {12657, 231}, {33726, 82}, {53652, 239}, {81410, 621}, {93968, 124}, {96560, 336}, {87170, 391}, {99709, 339}, {62624, 391}, {87539, 338}, {86453, 429}, {59868, 63}, {84429, 381}, {30486, 232}, {40567, 270}, {49919, 758}, {27419, 295}, {34130, 630}, {48132, 829}, {66593, 267}, {62012, 450}, {6766, 966}, {74049, 289}, {76198, 203}, {51774, 315}, {29504, 415}, {63400, 678}, {88801, 171}, {94208, 578}, {92122, 142}, {17728, 983}, {89983, 137}, {10780, 300}, {6589, 355}, {77263, 81}, {20119, 469}, {42538, 56}, {5626, 918}, {43470, 765}, {26363, 248}, {28816, 566}, {17649, 281}, {26681, 169}, {24915, 872}, {58031, 579}, {70752, 861}, {57072, 541}, {98391, 169}, {43778, 472}, {23142, 328}, {23502, 819}, {47644, 100}, {94031, 418}, {77959, 378}, {92389, 174}, {66922, 220}, {19838, 504}, {91208, 669}, {21547, 373}, {63374, 589}, {9915, 559}, {15602, 896}, {42375, 313}, {53438, 633}, {21624, 802}, {42772, 727}, {91058, 461}, {68824, 611}, {62817, 508}, {93455, 766}, {20345, 733}, {95223, 769}, {39557, 997}, {68535, 665}, {64855, 212}, {32494, 448}, {3796, 484}, {25111, 872}, {65088, 572}, {54428, 167}, {9872, 553}, {91550, 812}, {76039, 675}, {47025, 299}, {87388, 872}, {42779, 301}, {24375, 320}, {39294, 91}, {38634, 136}, {48286, 38}, {1963, 692}, {6695, 414}, {17804, 459}, {61632, 447}, {75107, 673}, {6829, 76}, {50617, 77}, {80115, 436}, {37780, 186}, {63222, 509}, {75456, 235}, {11951, 732}, {75266, 90}, {34824, 375}, {83918, 200}, {35354, 414}, {1317, 641}, {59904, 917}, {4957, 207}, {97110, 563}, {38963, 288}, {45036, 729}, {47325, 796}, {59891, 963}, {21204, 379}, {62647, 339}, {94888, 706}, {43560, 959}, {69989, 353}, {60477, 606}, {13058, 503}, {87661, 647}, {37988, 188}, {71262, 746}, {99289, 942}, {63876, 307}, {5733, 675}, {64252, 165}, {44518, 777}, {18278, 697}, {66308, 939}, {56356, 20}, {24162, 534}, {2350, 802}, {21729, 865}, {96255, 432}, {91396, 913}, {91505, 821}, {95068, 62}, {75451, 988}, {70513, 289}, {69448, 794}, {87550, 678}, {31982, 763}, {72954, 994}, {11032, 296}, {5822, 671}, {21413, 459}, {42375, 559}, {48136, 33}, {36419, 277}, {94437, 515}, {51972, 971}, {18225, 239}, {98319, 792}, {62859, 721}, {96850, 28}, {8015, 389}, {31863, 393}, {59191, 446}, {45943, 37}, {57191, 111}, {30477, 14}, {2616, 842}, {65180, 519}, {37177, 228}, {9355, 754}, {7345, 124}, {42035, 543}, {33708, 107}, {80678, 187}, {38172, 501}, {55328, 378}, {29332, 939}, {88254, 236}, {1144, 795}, {15472, 701}, {43559, 574}, {44058, 711}, {96432, 449}, {89939, 269}, {6999, 922}, {62239, 113}, {87631, 898}, {29496, 763}, {922, 124}, {39177, 467}, {93960, 71}, {55208, 840}, {74497, 279}, {83520, 840}, {47531, 970}, {2229, 818}, {96851, 670}, {64034, 21}, {75893, 336}, {10033, 547}, {86934, 368}, {81162, 911}, {82501, 289}, {74564, 662}, {63139, 162}, {92355, 393}, {68181, 249}, {80706, 391}, {60360, 321}, {32570, 738}, {35866, 253}, {57976, 380}, {37724, 974}, {89676, 306}, {68003, 903}, {5010, 990}, {43742, 204}, {77532, 890}, {31463, 427}, {44049, 893}, {3461, 759}, {27975, 126}, {75704, 113}, {98951, 310}, {7955, 977}, {87256, 873}, {31052, 961}, {29556, 982}, {90945, 11}, {15064, 509}, {64310, 399}, {69246, 853}, {3426, 230}, {27823, 103}, {25513, 447}, {17803, 267}, {58748, 4}, {34656, 757}, {64695, 482}, {98375, 852}, {88788, 664}, {50377, 710}, {38779, 540}, {48467, 194}, {8080, 670}, {23747, 347}, {82129, 108}, {39604, 240}, {29477, 530}, {29946, 602}, {70829, 958}, {56107, 723}, {13195, 525}, {31146, 139}, {19749, 894}, {13207, 790}, {98688, 992}, {79619, 812}, {39022, 7}, {16117, 367}, {79498, 964}, {72677, 609}, {98872, 342}, {53074, 700}, {62740, 673}, {47568, 862}, {84592, 399}, {36035, 533}, {83344, 915}, {90276, 796}, {58582, 92}, {53117, 757}, {87485, 816}, {41327, 82}, {5540, 350}, {79793, 855}, {17086, 403}, {36898, 599}, {98441, 430}, {94592, 844}, {14188, 594}, {15086, 556}, {22196, 392}, {70412, 683}, {7920, 229}, {59955, 937}, {36759, 680}, {11878, 441}, {17882, 185}, {41438, 813}, {42181, 917}, {1932, 837}, {17122, 357}, {87825, 502}, {75823, 235}, {84519, 312}, {58991, 956}, {17069, 154}, {29753, 941}, {54107, 848}, {79417, 548}, {29062, 177}, {60761, 490}, {5565, 707}, {89559, 480}, {3485, 571}, {88442, 463}, {13975, 5}, {42292, 60}, {58846, 132}, {64815, 491}, {21328, 801}, {53994, 602}, {98316, 467}, {44900, 477}, {90444, 980}, {9762, 780}, {82911, 739}, {25544, 769}, {81008, 491}, {62782, 103}, {62568, 483}, {36526, 543}, {69552, 800}, {20813, 966}, {82440, 629}, {34518, 540}, {66207, 959}, {68447, 560}, {96099, 213}, {25014, 624}, {66459, 887}, {22166, 776}, {18156, 266}, {71811, 23}, {29259, 199}, {47720, 436}, {49886, 719}, {32687, 72}, {56595, 326}, {59885, 120}, {46075, 783}, {24702, 202}, {34946, 941}, {61842, 174}, {37591, 434}, {7838, 924}, {55823, 325}, {35227, 111}, {15904, 490}, {61354, 291}, {51543, 399}, {5493, 424}, {27168, 184}, {63546, 350}, {90406, 659}, {55945, 239}, {36663, 639}, {21228, 169}, {22540, 636}, {51317, 228}, {31331, 785}, {90173, 953}, {52667, 330}, {3618, 511}, {18151, 258}, {51434, 567}, {24318, 700}, {82803, 850}, {80399, 646}, {21118, 398}, {88200, 299}, {64160, 266}, {38559, 896}, {42892, 344}, {24340, 116}, {72413, 205}, {10818, 953}, {19068, 105}, {43797, 439}, {23462, 951}, {25968, 293}, {79530, 483}, {72085, 591}, {83341, 699}, {43724, 905}, {72218, 285}, {79563, 791}, {98626, 629}, {50159, 510}, {47172, 752}, {43412, 399}, {78403, 965}, {38480, 901}, {43992, 799}, {79932, 611}, {56912, 391}, {22010, 304}, {58096, 610}, {5060, 72}, {61726, 444}, {95732, 831}, {42070, 530}, {87299, 466}, {89369, 637}, {88436, 574}, {75067, 478}, {66237, 593}, {12161, 718}, {8902, 75}, {91909, 49}, {54481, 166}, {42032, 254}, {69290, 951}, {87879, 932}, {22476, 326}, {37817, 318}, {7294, 626}, {63918, 984}, {19183, 931}, {26364, 154}, {12181, 33}, {78103, 388}, {2090, 975}, {3413, 697}, {14007, 114}, {22462, 89}, {34884, 496}, {26863, 645}, {26845, 267}, {67018, 708}, {38576, 639}, {80011, 455}, {7457, 723}, {80991, 880}, {64984, 997}, {20518, 267}, {39859, 876}, {13782, 29}, {72459, 765}, {63813, 32}, {44700, 271}, {55979, 738}, {56095, 596}, {5045, 705}, {75495, 482}, {84657, 834}, {2214, 124}, {9712, 701}, {72427, 112}, {56139, 27}, {63451, 532}, {54093, 742}, {39819, 101}, {52558, 605}, {47314, 240}, {71349, 807}, {26789, 437}, {67052, 198}, {29211, 620}, {22086, 578}, {14650, 775}, {97152, 438}, {92949, 11}, {20436, 777}, {15934, 982}, {27990, 104}, {95072, 114}, {46098, 481}, {21849, 391}, {25964, 627}, {47503, 409}, {35667, 113}, {34166, 565}, {21829, 228}, {94249, 391}, {12067, 606}, {528, 995}, {9006, 830}, {58025, 658}, {57320, 711}, {62239, 121}, {51316, 303}, {47424, 419}, {83001, 655}, {78315, 689}, {87290, 324}, {13083, 117}, {40516, 582}, {63738, 23}, {65536, 46}, {30775, 594}, {60066, 502}, {55692, 860}, {21946, 577}, {63195, 481}, {70751, 805}, {88831, 678}, {13384, 951}, {87911, 56}, {95117, 558}, {88995, 575}, {91456, 888}, {35598, 972}, {84956, 466}, {12234, 820}, {65858, 757}, {4182, 210}, {26918, 114}, {51970, 380}, {99227, 51}, {70275, 607}, {62739, 801}, {86666, 80}, {44486, 640}, {33796, 923}, {55014, 228}, {10317, 222}, {96070, 85}, {40263, 100}, {86160, 114}, {41735, 857}, {56607, 500}, {45452, 488}, {3203, 276}, {60684, 776}, {45188, 952}, {7832, 113}, {76240, 138}, {89832, 504}, {85047, 88}, {93939, 894}, {14182, 363}, {89298, 32}, {77604, 724}, {53187, 233}, {99436, 186}, {25190, 564}, {35006, 370}, {39571, 988}, {29207, 783}, {13816, 253}, {99485, 191}, {32274, 980}, {22502, 160}, {4112, 725}, {54296, 649}, {17155, 684}, {62383, 764}, {37144, 292}, {92367, 257}, {6622, 136}, {2216, 25}, {30617, 735}, {21698, 638}, {52246, 874}, {60939, 137}, {61120, 716}, {42730, 315}, {85293, 178}, {78636, 512}, {7594, 871}, {98250, 424}, {69204, 672}, {79431, 737}, {26956, 269}, {90320, 366}, {98948, 337}, {95329, 143}, {40206, 256}, {46984, 170}, {57383, 804}, {79601, 310}, {27020, 77}, {45029, 90}, {64413, 87}, {90569, 319}, {43208, 422}, {68443, 524}, {74629, 606}, {64614, 908}, {26018, 876}, {67060, 665}, {24668, 90}, {65334, 562}, {3692, 383}, {57224, 268}, {49635, 511}, {37583, 814}, {18623, 895}, {84371, 628}, {91786, 998}, {19600, 744}, {35792, 743}, {58414, 286}, {49416, 962}, {90190, 712}, {90045, 337}, {13547, 357}, {21212, 363}, {68115, 16}, {37100, 293}, {27144, 455}, {54884, 80}, {41199, 238}, {69856, 826}, {52338, 347}, {67154, 162}, {30547, 993}, {6504, 911}, {70240, 369}, {21788, 376}, {83944, 263}, {78171, 918}, {2647, 514}, {88382, 333}, {6111, 114}, {93158, 587}, {37613, 551}, {153, 377}, {73669, 784}, {69533, 274}, {12866, 502}, {85986, 670}, {8158, 433}, {5693, 107}, {31457, 863}, {65595, 992}, {65, 832}, {95294, 267}, {43418, 834}, {82945, 351}, {58734, 670}, {28973, 716}, {50186, 637}, {89384, 788}, {44110, 90}, {82700, 608}, {20759, 387}, {63640, 128}, {97033, 982}, {77222, 834}, {38366, 729}, {40275, 718}, {81895, 165}, {89881, 682}, {59148, 884}, {59648, 9}, {70967, 585}, {96854, 329}, {58613, 326}, {78002, 640}, {80164, 935}, {84420, 663}, {61442, 372}, {71586, 237}, {83866, 981}, {51978, 654}, {80372, 675}, {73738, 980}, {88032, 577}, {65866, 794}, {14177, 772}, {75270, 23}, {91533, 108}, {71978, 647}, {98978, 348}, {80598, 489}, {48645, 466}, {55816, 283}, {53766, 29}, {38501, 430}, {4512, 264}, {41780, 208}, {83863, 750}, {90098, 847}, {66802, 330}, {43570, 990}, {72218, 499}, {20349, 787}, {10658, 125}, {5952, 353}, {32269, 819}, {5450, 721}, {26293, 690}, {17338, 428}, {52276, 320}, {20948, 302}, {2230, 655}, {10225, 443}, {97926, 22}, {77700, 693}, {69157, 838}, {2371, 609}, {70074, 4}, {73354, 594}, {28224, 631}, {89406, 296}, {35389, 925}, {36489, 701}, {52860, 137}, {10336, 410}, {50829, 385}, {38373, 743}, {82074, 545}, {53828, 920}, {61381, 995}, {39323, 356}, {77276, 156}, {73328, 229}, {51754, 657}, {87008, 858}, {94974, 748}, {6495, 862}, {50524, 994}, {55340, 240}, {26463, 646}, {56604, 906}, {35118, 749}, {5107, 993}, {51053, 849}, {96753, 379}, {63992, 505}, {11377, 20}, {92750, 805}, {80594, 612}, {58747, 295}, {5325, 98}, {27907, 848}, {76483, 400}, {831, 789}, {77805, 267}, {89628, 546}, {58099, 353}, {15603, 265}, {63330, 701}, {10020, 734}, {92598, 126}, {57601, 722}, {86348, 208}, {3968, 68}, {5740, 48}, {59508, 124}, {47483, 359}, {42520, 413}, {18553, 710}, {11233, 129}, {30822, 380}, {20782, 814}, {69471, 740}, {91531, 885}, {14961, 668}, {94530, 978}, {31449, 245}, {84330, 616}, {15122, 434}, {35438, 223}, {3495, 224}, {31045, 683}, {46023, 602}, {36371, 568}, {85210, 525}, {5028, 749}, {61874, 376}, {94820, 373}, {66905, 542}, {96602, 906}, {14609, 102}, {56030, 520}, {87645, 700}, {81427, 976}, {92606, 705}, {4393, 964}, {34917, 981}, {50433, 130}, {81884, 28}, {27316, 837}, {94809, 906}, {28518, 846}, {30480, 402}, {46027, 166}, {60750, 921}, {50682, 156}, {56964, 96}, {31062, 125}, {94508, 854}, {49723, 68}, {81314, 16}, {11973, 664}, {30310, 881}, {31051, 282}, {63272, 981}, {70649, 694}, {61477, 27}, {52696, 508}, {8411, 953}, {74638, 758}, {66281, 433}, {36558, 2}, {51386, 164}, {47946, 88}, {44160, 287}, {48306, 151}, {76279, 323}, {80689, 99}, {53114, 51}, {25726, 919}, {89534, 222}, {82013, 971}, {39050, 911}, {70705, 313}, {19463, 583}, {37382, 398}, {95478, 592}, {36847, 631}, {94200, 510}, {33174, 84}, {50564, 161}, {64832, 661}, {50112, 145}, {98613, 569}, {35253, 955}, {416, 685}, {13011, 675}, {61379, 125}, {95428, 706}, {86700, 518}, {45974, 223}, {80565, 664}, {50790, 247}, {81260, 388}, {90943, 441}, {9986, 627}, {52141, 290}, {72113, 803}, {30423, 996}, {42558, 41}, {17656, 48}, {79424, 407}, {1541, 388}, {77999, 870}, {97502, 574}, {12020, 277}, {57259, 982}, {73279, 322}, {81511, 253}, {16243, 511}, {88549, 117}, {71070, 863}, {69671, 888}, {56585, 906}, {3392, 535}, {15045, 194}, {62215, 765}, {52984, 123}, {74420, 421}, {52058, 608}, {24919, 169}, {10055, 539}, {59281, 211}, {48568, 577}, {3029, 415}, {42433, 851}, {38279, 897}, {65275, 885}, {1036, 158}, {92096, 887}, {20485, 393}, {52172, 473}, {35289, 676}, {25951, 426}, {168, 166}, {69777, 633}, {39874, 694}, {34935, 751}, {85593, 783}, {93526, 100}, {81060, 704}, {2447, 909}, {23119, 409}, {35117, 96}, {70310, 368}, {23096, 416}, {68078, 350}, {12749, 515}, {12963, 752}, {93778, 484}, {40299, 504}, {71714, 571}, {36888, 192}, {95755, 647}, {82923, 499}, {914, 644}, {18215, 16}, {35248, 249}, {46608, 401}, {55956, 354}, {28672, 371}, {32638, 534}, {22448, 341}, {34971, 178}, {63999, 247}, {94938, 359}, {51603, 844}, {85449, 68}, {31699, 535}, {11262, 154}, {76843, 277}, {96610, 554}, {96879, 960}, {34917, 150}, {23306, 160}, {15009, 197}, {6242, 320}, {32250, 687}, {35205, 320}, {92063, 119}, {34173, 870}, {9906, 758}, {92057, 199}, {32829, 668}, {64137, 642}, {29297, 577}, {76832, 325}, {11452, 234}, {98944, 866}, {48812, 102}, {28790, 471}, {93182, 535}, {97174, 235}, {88107, 674}, {93454, 67}, {46421, 739}, {33019, 568}, {36991, 134}, {97822, 673}, {22516, 252}, {46733, 48}, {38859, 384}, {61560, 644}, {81336, 198}, {74031, 518}, {78433, 439}, {72586, 211}, {41803, 362}, {34618, 441}, {68436, 597}, {53741, 962}, {63661, 120}, {40893, 939}, {20237, 462}, {42382, 110}, {59196, 70}, {21902, 459}, {36007, 995}, {35950, 819}, {12288, 378}, {74753, 169}, {66294, 208}, {37684, 307}, {81587, 979}, {4294, 897}, {90616, 960}, {59236, 62}, {66319, 831}, {72386, 850}, {72641, 845}, {27971, 16}, {409, 155}, {32779, 620}, {3165, 278}, {40678, 518}, {96415, 780}, {70806, 38}, {76134, 420}, {6027, 268}, {79770, 883}, {89200, 93}, {92051, 182}, {49847, 472}, {20848, 897}, {12258, 548}, {84580, 725}, {81435, 842}, {81917, 879}, {80750, 922}, {42719, 670}, {39416, 623}, {95502, 760}, {32016, 566}, {85776, 210}, {45929, 852}, {53794, 608}, {671, 256}, {5206, 723}, {40489, 828}, {81563, 391}, {29021, 477}, {56136, 105}, {67111, 500}, {45794, 261}, {80413, 723}, {75589, 761}, {53995, 712}, {58787, 772}, {53803, 382}, {40833, 701}, {54189, 627}, {3145, 573}, {40406, 985}, {90050, 404}, {67492, 698}, {70917, 384}, {69045, 273}, {74480, 199}}, 84113, 306)

	// 13357376
	// 3305058

	// minimumFinishTime([][]int{{1, 2}}, 10000, 100000)
}

func minimumFinishTimeBottomUp(tires [][]int, changeTime int, numLaps int) int {
	// type Tire struct {
	// 	F, R int
	// }
	// unique := make(map[Tire]struct{})
	// for _, v := range tires {
	// 	F, R := v[0], v[1]
	// 	unique[Tire{F, R}] = struct{}{}
	// }

	unique := make(map[int]struct{})
	for _, v := range tires {
		unique[v[0]*100001+v[1]] = struct{}{}
	}

	xs := make([][18]int, len(unique))

	i := 0
	for k := range unique {
		xs[i] = [18]int{}
		F, R := k/100001, k%100001
		// F, R := k.F, k.R

		xs[i][0] = F
		prev := F
		for lap := 1; lap < 18; lap++ {
			prev = prev * R
			xs[i][lap] = xs[i][lap-1] + prev
			if xs[i][lap] > changeTime {
				break
			}
		}
		i++
	}
	fmt.Println(xs)

	dp = make([]int, numLaps+1)
	min := math.MaxInt
	for _, v := range xs {
		if v[0] < min {
			min = v[0]
		}
	}
	dp[0] = -changeTime

	for i := 1; i < numLaps+1; i++ {
		dp[i] = math.MaxInt

		for idx, v := range xs {
			for j := 0; j < Min(i, 18) && v[j] != 0 && (i-j) >= 0; j++ {
				c := v[j] + changeTime + dp[i-j-1]
				fmt.Println("numlap", i, "tire", idx, j, "cost", v[j], changeTime, dp[i-j-1], "total_cost", c)
				dp[i] = Min(dp[i], c)
			}
		}
	}

	fmt.Println(dp[numLaps])

	return dp[numLaps]
}

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	type Tire struct {
		F, R int
	}
	unique := make(map[Tire]struct{})
	for _, v := range tires {
		F, R := v[0], v[1]
		unique[Tire{F, R}] = struct{}{}
	}

	xs := make([][18]int, len(unique))

	i := 0
	for k := range unique {
		xs[i] = [18]int{}

		// F, R := v[0], v[1]
		F, R := k.F, k.R

		xs[i][0] = F
		prev := F
		for lap := 1; lap < 18; lap++ {
			prev = prev * R
			xs[i][lap] = xs[i][lap-1] + prev
			if xs[i][lap] > changeTime {
				break
			}
		}
		i++
	}
	fmt.Println(xs)

	dp = make([]int, numLaps+1)

	cost := MinLapTime(xs, changeTime, numLaps)
	fmt.Println("total", cost)
	return 0
}

// [70917 27232128 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [69045 18849285 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [74480 14821520 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
// [70917 27232128 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [69045 18849285 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [74480 14821520 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

var dp []int

func MinLapTime(tires [][18]int, changeTime int, remaining int) int {
	if remaining <= 0 {
		return -changeTime
	}

	if dp[remaining] != 0 {
		return dp[remaining]
	}

	cost := math.MaxInt
	for _, v := range tires {
		// cc := 0
		for i := 0; i < Min(remaining, 18); i++ {
			if v[i] == 0 {
				break
			}
			// fmt.Println("remaining", remaining, "checking for", tire, "using", i, v[i])

			// cc = cc + v[i]

			c := MinLapTime(tires, changeTime, remaining-i-1) + v[i] + changeTime
			cost = Min(cost, c)
		}
	}
	// fmt.Println("remaining", remaining, cost)
	dp[remaining] = cost

	return cost
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// type Tire struct {
// 	F, R, X     int
// 	ChangeCost  int
// 	CurrentCost int
// 	MaxLap      int
// }

// func (t Tire) Cost() int {
// 	// Cost = f * r^(x-1)
// 	return t.F * Pow(t.R, t.X-1)
// }

// func (t Tire) Increment() int {
// 	t.CurrentCost *= t.R
// 	return t.CurrentCost
// }

// func CalculateCost(F, R, X int) int {
// 	return F * Pow(R, X-1)
// }

// func Pow(n, m int) int {
// 	if m == 0 {
// 		return 1
// 	}
// 	if m == 1 {
// 		return n
// 	}

// 	r := n
// 	for i := 2; i <= m; i++ {
// 		r *= n
// 	}
// 	return r
// }

// func Min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// // var dp []map[int]int
// var dp [][]int

// func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
// 	xs := make([]Tire, len(tires))
// 	for i, v := range tires {
// 		F, R := v[0], v[1]
// 		cost := CalculateCost(F, R, 1)
// 		xs[i] = Tire{
// 			F:           F,
// 			R:           R,
// 			X:           1,
// 			ChangeCost:  cost + changeTime,
// 			CurrentCost: cost,
// 			MaxLap:      1,
// 		}
// 	}

// 	for i, v := range xs {
// 		for v.CurrentCost < v.ChangeCost {
// 			xs[i].MaxLap++
// 			v.CurrentCost *= v.R
// 		}
// 	}

// 	fmt.Println(xs)

// 	// fmt.Println(xs)
// 	// dp = make([]map[int]int, numLaps+1)
// 	// for i := 0; i < numLaps+1; i++ {
// 	// 	dp[i] = make(map[int]int)
// 	// }

// 	dp = make([][]int, numLaps+1)
// 	for i := 0; i < numLaps+1; i++ {
// 		dp[i] = make([]int, 1000000)
// 		for j := range dp[i] {
// 			dp[i][j] = -1
// 		}
// 	}

// 	cost := math.MaxInt
// 	// for i, v := range xs {
// 	// 	// v := v

// 	// 	// cc := v.Cost()
// 	// 	// v.X++

// 	// 	c := minTime(xs, v, 1, i, numLaps) + v.CurrentCost

// 	// 	cost = Min(cost, c)
// 	// }

// 	// cost := minTime(xs, Tire{10000, 10000, 2}, changeTime, 0, -1, numLaps)

// 	// fmt.Println(dp)
// 	// fmt.Println("total cost", cost)
// 	return cost
// }
// func minTime(tires []Tire, tire Tire, numLaps int, idx int, maxLap int) int {
// 	// 2 option
// 	// change tire
// 	// using current tire

// 	// fmt.Println("dp", numLaps, "tire", tire, tire.Cost(), "idx", idx, tire.X)

// 	if numLaps == maxLap {
// 		return 0
// 	}

// 	// fmt.Println("from prev tire", tire, idx, "with cost", old)

// 	// dpKey := idx*10000 + tire.X
// 	// memo, ok := dp[numLaps][dpKey]
// 	// if ok {
// 	// 	return memo
// 	// }

// 	dpKey := idx*1000 + tire.X
// 	// memo := dp[numLaps][dpKey]
// 	// if memo != 1 {
// 	// 	return memo
// 	// }

// 	var shouldUseNewTire bool
// 	// tire.X++
// 	// for _, v := range tires {
// 	// 	// c := v.CurrentCost + changeTime
// 	// 	if v.ChangeCost < tire.CurrentCost {
// 	// 		shouldUseNewTire = true
// 	// 		break
// 	// 	}
// 	// }
// 	// tire.X--

// 	// fmt.Println(shouldUseNewTire, idx)

// 	cost := math.MaxInt

// 	for i, v := range tires {
// 		c := minTime(tires, v, numLaps+1, i, maxLap) + v.ChangeCost
// 		cost = Min(cost, c)

// 		if v.ChangeCost < tire.CurrentCost {
// 			shouldUseNewTire = true
// 		}
// 	}

// 	if !shouldUseNewTire {
// 		tire.X++
// 		tire.CurrentCost = tire.CurrentCost * tire.R
// 		c := minTime(tires, tire, numLaps+1, idx, maxLap) + tire.CurrentCost // + tire.Increment()
// 		cost = Min(cost, c)
// 	}

// 	dp[numLaps][dpKey] = cost

// 	// fmt.Println(numLaps, "pick", idx, cost)
// 	return cost
// }
