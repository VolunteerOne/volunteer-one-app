import React, {useState} from "react";
import { StyleSheet, Dimensions, ScrollView, Text, View } from "react-native";
import { Block, theme } from "galio-framework";
import NewPostModal from "../../components/Modals/NewPostModal";
import profiles from "../../constants/ProfileTab/profile";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";
import { Button } from "../../components";
import argonTheme from "../../constants/Theme";
import { PostImageCard, PostNoImageCard } from "../../components";
import mockPosts from "../../constants/FeedTab/posts";

// ================================= View Feed Page ================================= //

const { width } = Dimensions.get("screen");

// class Feed extends React.Component {
const Feed = ({ navigation, route }) => {
  const JESSICA = "Jessica Jones";
  const [user, setUser] = useState(profiles[JESSICA]);

  const [modalVisible, setModalVisible] = useState(false);
  const [posts, setPosts] = useState(mockPosts);

  const handleModalVisible = () => {
    setModalVisible(!modalVisible);
  };

  const addNewPost = (data) => {
    // console.log("Adding New Post", data);
    const newPost = {
      id: posts.length + 1,
      author: user.name,
      timePosted: data["datetime"],
      profileImage: user.image,
      description: data["description"],
      image: data["image"],
    };
    // console.log(newPost);
    setPosts([newPost, ...posts]);
  };

  renderPosts = () => {
    var postsList = posts.map((data) => {
        if (data["image"] != null && data.id != null)
          return (
            <View key={data.id}>
            <PostImageCard key={data.id} data={data} />
            </View>
          );
        else if (data.id != null)
          return (
            <View key={data.id}>
            <PostNoImageCard key={data.id} data={data} />
            </View>
          );
    });

    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.posts}
      >
        <Block flex center>
          {postsList}
        </Block>
      </ScrollView>
    );
  };

  return (
    <Block flex center style={styles.home}>
      <Block middle>
        <Button
          color="primary"
          style={styles.button}
          onPress={() => handleModalVisible()}
        >
          <Block row middle>
            <MaterialCommunityIcons
              size={24}
              name="plus-box-outline"
              color={theme.COLORS.WHITE}
            />
            <Text bold size={14} style={styles.buttonTitle}>
              New Post
            </Text>
          </Block>
        </Button>
      </Block>
      {modalVisible && (
        <NewPostModal
          visible={modalVisible}
          handleModalVisible={handleModalVisible}
          addNewPost={addNewPost}
        />
      )}

      {renderPosts()}
    </Block>
  );
};

const styles = StyleSheet.create({
  home: {
    width: width,
  },
  posts: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
  button: {
    marginTop: theme.SIZES.BASE,
    marginBottom: 0,
    width: width * 0.9,
  },
  buttonTitle: {
    paddingLeft: 5,
    lineHeight: 19,
    fontWeight: "600",
    color: argonTheme.COLORS.WHITE,
  },
});

export default Feed;
