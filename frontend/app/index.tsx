import { ActivityIndicator, Button, StyleSheet, View } from "react-native";
import { AuthProvider, useAuth } from "./context/AuthContext";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import Login from "./screens/login";
import Signup from "./screens/signup";
import Profile from "./screens/profile";
import { RootStackParamList } from "./navigarionTypes";
import { style } from "./screens/styles/root";

const Stack = createNativeStackNavigator<RootStackParamList>();

const linking = {
  prefixes: ["yourapp://", "https://yourapp.com"],
  config: {
    screens: {
      Login: "login",
      Signup: "signup",
      Profile: "profile",
    },
  },
};

export default function App() {
  return (
    <AuthProvider>
      <Layout />
    </AuthProvider>
  );
}

const Layout = () => {
  const { authState, onLogout } = useAuth();

  if (authState?.authenticated === null) {
    return (
      <View style={style.container}>
        <ActivityIndicator size="large" color="#0000ff" />
      </View>
    );
  }
  return (
    <Stack.Navigator>
      {authState?.authenticated ? (
        <Stack.Screen
          name="Profile"
          component={Profile}
          options={{
            headerRight: () => (
              <View style={{ marginRight: 10 }}>
                <Button title="Logout" onPress={onLogout} />
              </View>
            ),
          }}
        />
      ) : (
        <>
          <Stack.Screen name="Login" component={Login} />
          <Stack.Screen
            name="Signup"
            options={{ title: "Sign up" }}
            component={Signup}
          />
        </>
      )}
    </Stack.Navigator>
  );
};
