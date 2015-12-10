// Day 2 of adventofcode.com

def int calculatePaperForPackage(int length, int width, int height){
	int size, min, currentArea = 0
	currentArea = length * width;
	min = currentArea;
	size += 2* currentArea;
	currentArea = width * height;
	if(currentArea < min) min = currentArea;
	size += 2* currentArea;
	currentArea = height * length;
	if(currentArea < min) min = currentArea;
	size += 2* currentArea;
	return size + min;
}

def int calculateRibbonLength(int length, int width, int height){
	int bow = length * width * height;
	int ribbonlength = 0;
	int[] mins = new int[2];
	def swap = {
		array, one, two -> 
			int swap = array[one];
			array[one] = array[two];
			array[two] = swap;
	}
	int[] lengths = [length, width, height];
	if(lengths[0] > lengths[2]){
		swap(lengths, 0, 2);
	}
	if(lengths[1] < lengths[0]){
		swap(lengths, 1, 0);
	} else if (lengths[1] > lengths[2]){
		swap(lengths, 1, 2);
	}
	ribbonlength = 2*lengths[0] + 2*lengths[1] + bow;
}

File file = new File('input.txt')
int completeAmountOfPaper = 0;
int ribbonlength = 0;
if (file.length() > 0){
	file.eachLine {
	line ->	
		values = line.split('x');
		completeAmountOfPaper += calculatePaperForPackage(new Integer(values[0]), new Integer(values[1]), new Integer(values[2]));
		ribbonlength += calculateRibbonLength(new Integer(values[0]), new Integer(values[1]), new Integer(values[2]));
}

println("Amount needed: "+completeAmountOfPaper )
print("Ribbon needed: "+ribbonlength)
}
