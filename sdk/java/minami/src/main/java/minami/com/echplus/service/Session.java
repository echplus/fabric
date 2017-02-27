package minami.com.echplus.service;

import java.util.Properties;

import javax.annotation.PostConstruct;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.springframework.stereotype.Service;

import net.sf.ehcache.CacheManager;
import net.sf.ehcache.Ehcache;
import net.sf.ehcache.Element;

@Service
public class Session {

    private static final Logger logger = LogManager.getLogger(Session.class);

    private static int timeToLiveSeconds = 600;

    private static final CacheManager cacheManager = new CacheManager();

    private static final Ehcache session = cacheManager.addCacheIfAbsent("session");

    private static final Properties properties = new Properties();

    @PostConstruct
    private void init() throws Exception {
        properties.load(Session.class.getClassLoader().getResourceAsStream("/application.properties"));
        timeToLiveSeconds = Integer.parseInt(properties.getProperty("timeToLiveSeconds"));
    }

    public int GeTimeToLiveSeconds() {
        return timeToLiveSeconds;
    }

    public boolean Exists(String token) {
        return session.get(token) != null;
    }

    public void Update(String token) {
        Element tk = session.get(token);
        if (tk == null) {
            tk = new Element(token, token);
        }
        tk.setTimeToLive(timeToLiveSeconds);
        session.put(tk);
    }

    @SuppressWarnings("unused")
    private void Print() {
        for (Object key : session.getKeys()) {
            logger.info(key.toString() + "  " + session.get(key).getTimeToLive());
        }
        logger.info("--------------------------------");
    }

}
