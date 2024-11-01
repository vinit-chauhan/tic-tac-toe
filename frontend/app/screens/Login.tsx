import { useState } from "react";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from "react-native";
import { useAuth } from "../context/AuthContext";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const { onLogin } = useAuth();

  const handleLogin = async () => {
    const result = await onLogin!(username, password);
    if (result && result.error) {
      alert(result.msg);
    }
  };

  return (
    <SafeAreaView style={[style.container, style.backdrop]}>
      <View style={[style.container, style.formGroup]}>
        <TextInput
          style={style.textBox}
          placeholder="Username"
          onChangeText={(username) => setUsername(username)}
        />

        <TextInput
          style={style.textBox}
          placeholder="Password"
          autoComplete="current-password"
          secureTextEntry={true}
          onChangeText={(password) => setPassword(password)}
        />

        <TouchableOpacity onPress={handleLogin} style={[style.loginBtn]}>
          <Text>Sign in</Text>
        </TouchableOpacity>

        <TouchableOpacity onPress={() => {}}>
          <Text style={style.link}>Don't have an account?</Text>
        </TouchableOpacity>
      </View>
    </SafeAreaView>
  );
};

export default Login;

const style = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
  textBox: {
    borderRadius: 8,
    borderColor: "#F35369AA",
    borderWidth: 2,
    height: 50,
    width: "80%",
    justifyContent: "center",
    alignItems: "center",
    padding: 10,
    margin: 10,
  },
  link: {
    margin: 15,
    textAlign: "center",
    color: "#00f",
  },
  loginBtn: {
    width: "40%",
    borderRadius: 8,
    height: 50,
    alignItems: "center",
    justifyContent: "center",
    marginTop: 10,
    backgroundColor: "#F35369",
  },
  backdrop: {
    backgroundColor: "#FAFFD2",
  },
  formGroup: {
    width: "50%",
  },
});
