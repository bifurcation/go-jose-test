import org.jose4j.jwk.*;
import org.jose4j.jws.*;
import java.security.*;

// javac -classpath jose4j-0.4.2.jar JwsTest.java
// java -classpath jose4j-0.4.2.jar:. JwsTest

class JwsTest {
  public static void main(String[] args) {
    try {
      KeyPairGenerator keyGen = KeyPairGenerator.getInstance("RSA");
      keyGen.initialize(2048);
      KeyPair keyPair = keyGen.genKeyPair();
      JsonWebKey jwk = JsonWebKey.Factory.newJwk(keyPair.getPublic());

      String payload = "This is some text that is to be signed.";

      JsonWebSignature jws = new JsonWebSignature();
      jws.setPayload(payload);
      jws.setAlgorithmHeaderValue(AlgorithmIdentifiers.RSA_USING_SHA256);
      jws.setKey(keyPair.getPrivate());
      jws.getHeaders().setJwkHeaderValue("jwk", jwk);

      System.out.println(jws.getCompactSerialization());
    } catch (Exception e) {
      // XXX
      e.printStackTrace();
    }
  }
}
