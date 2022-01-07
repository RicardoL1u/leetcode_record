#include<iostream>
using namespace std;
int findK(int *nums, int k, int start, int end){
    int low = start;
    int high = end;
    int temp = nums[low];
    cout << high << endl;
    while(low < high){
        while(low < high && nums[high] <= temp){
            high --;
        } 
        nums[low] = nums[high];
        while(low<high && nums[low] >= temp){
            low++;
        }
        nums[high] = nums[low];
    }
    nums[high] = temp;
    if(high == k - 1){
        return temp;
    } else if (high > k - 1){
        return findK(nums,k,start,high - 1);
    } else {
        return findK(nums,k,high + 1,end);
    }
}

int main(){
    int nums[9] = {1,3,4,5,7,6,2,8,9};
    
}