package game

// The i'th element of these arrays is the bitboard for the moves available for a piece located at Square(i)
var (
	BBRookMoves = [36]BitBoard{
		1090785406, 2181570749, 4363141435, 8726282807, 17452565551, 34905131039,
		1090789249, 2181574466, 4363144900, 8726285768, 17452567504, 34905130976,
		1091035201, 2181812354, 4363366660, 8726475272, 17452692496, 34905126944,
		1106776129, 2197037186, 4377559300, 8738603528, 17460691984, 34904868896,
		2114195521, 3171426434, 5285888260, 9514811912, 17972659216, 34888353824,
		66589036609, 65532338306, 63418941700, 59192148488, 50738562064, 33831389216,
	}

	BBKingMoves = [36]BitBoard{
		194, 453, 906, 1812, 3624, 3088,
		12419, 28999, 57998, 115996, 231992, 197680,
		794816, 1855936, 3711872, 7423744, 14847488, 12651520,
		50868224, 118779904, 237559808, 475119616, 950239232, 809697280,
		3255566336, 7601913856, 15203827712, 30407655424, 60815310848, 51820625920,
		2197815296, 5486149632, 10972299264, 21944598528, 43889197056, 17985175552,
	}

	BBKnightMoves = [36]BitBoard{
		8448, 20992, 42048, 84096, 164096, 66048,
		540676, 1343496, 2691089, 5382178, 10502148, 4227080,
		34603266, 85983749, 172229706, 344459412, 672137512, 270533136,
		2214609024, 5502959936, 11022701184, 22045402368, 43016800768, 17314120704,
		4296024064, 8592052224, 18258108416, 36516216832, 4296179712, 8592097280,
		67633152, 135528448, 287834112, 575668224, 77594624, 138412032,
	}

	BBQueenMoves = [36]BitBoard{
		35721072894, 2722669053, 4371600315, 8726685495, 17469885999, 36013509663,
		18405932931, 36811865927, 4904513230, 8752057820, 18561076216, 37121886192,
		9748603077, 19497210314, 39010939797, 10375886634, 19677773332, 39338507304,
		5435306313, 10870878866, 22798984548, 45581453961, 22426912018, 43763304996,
		4262220369, 8541480098, 16025999620, 30995030600, 60932834449, 52072450338,
		66623673441, 65618389122, 63591301380, 59536609800, 51410707536, 34101938337,
	}

	BBWhitePawnPushes = [36]BitBoard{
		64, 128, 256, 512, 1024, 2048,
		4096, 8192, 16384, 32768, 65536, 131072,
		262144, 524288, 1048576, 2097152, 4194304, 8388608,
		16777216, 33554432, 67108864, 134217728, 268435456, 536870912,
		1073741824, 2147483648, 4294967296, 8589934592, 17179869184, 34359738368,
		0, 0, 0, 0, 0, 0,
	}

	BBWhitePawnCaptures = [36]BitBoard{
		128, 320, 640, 1280, 2560, 1024,
		8192, 20480, 40960, 81920, 163840, 65536,
		524288, 1310720, 2621440, 5242880, 10485760, 4194304,
		33554432, 83886080, 167772160, 335544320, 671088640, 268435456,
		2147483648, 5368709120, 10737418240, 21474836480, 42949672960, 17179869184,
		0, 0, 0, 0, 0, 0,
	}

	BBBlackPawnPushes = [36]BitBoard{
		0, 0, 0, 0, 0, 0,
		1, 2, 4, 8, 16, 32,
		64, 128, 256, 512, 1024, 2048,
		4096, 8192, 16384, 32768, 65536, 131072,
		262144, 524288, 1048576, 2097152, 4194304, 8388608,
		16777216, 33554432, 67108864, 134217728, 268435456, 536870912,
	}

	BBBlackPawnCaptures = [36]BitBoard{
		0, 0, 0, 0, 0, 0,
		2, 5, 10, 20, 40, 16,
		128, 320, 640, 1280, 2560, 1024,
		8192, 20480, 40960, 81920, 163840, 65536,
		524288, 1310720, 2621440, 5242880, 10485760, 4194304,
		33554432, 83886080, 167772160, 335544320, 671088640, 268435456,
	}
)