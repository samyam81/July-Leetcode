import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Stack;

public class RobotCollison {
    public List<Integer> survivedRobotsHealths(int[] positions, int[] healths, String directions) {
        List<Integer> list = new ArrayList<>();

        for (int i = 0; i < positions.length; i++) {
            list.add(i);
        }

        Collections.sort(list, (a, b) -> Integer.compare(positions[a], positions[b]));
        
        Stack<Integer> stack = new Stack<>();
        
        for (var l : list) {
            if (directions.charAt(l) == 'L') {
                while (!stack.isEmpty()) {
                    int temp = stack.peek();
                    
                    if (healths[l] == healths[temp]) {
                        healths[l] = 0;
                        healths[temp] = 0;
                        stack.pop();
                        break;
                    } else if (healths[l] > healths[temp]) {
                        healths[l]--;
                        healths[temp] = 0;
                        stack.pop();
                    } else {
                        healths[l] = 0;
                        healths[temp]--;
                        break;
                    }
                }
            } else {
                stack.push(l);
            }
        }
        
        List<Integer> res = new ArrayList<>();
        for (var h : healths) {
            if (h != 0)
                res.add(h);
        }
        
        return res;
    }
}