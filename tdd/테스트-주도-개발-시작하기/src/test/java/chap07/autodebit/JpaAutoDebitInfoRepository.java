package chap07.autodebit;

import java.util.HashMap;
import java.util.Map;

public class JpaAutoDebitInfoRepository implements AutoDebitInfoRepository {
    public Map<String, AutoDebitInfo> infos = new HashMap<>();

    @Override
    public void save(AutoDebitInfo info) {
        infos.put(info.getUserId(), info);
    }

    @Override
    public AutoDebitInfo findOne(String userId) {
        return infos.get(userId);
    }

}
