// This page is temporarily being used to test other pages (notifications, friends, etc.) - matt
import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
import { Block, theme } from 'galio-framework';

// ================================= View Friends Page ================================= //
import { PostImageCard, PostNoImageCard } from '../../components';
import posts from '../../constants/posts';

const { width } = Dimensions.get('screen');

class Friends extends React.Component {
  renderPosts = () => {
    var postsList = posts.map(function (data) {
      if (data["image"] == null)
        return <PostNoImageCard key={data["id"]} data={data} />;
      else
        return <PostImageCard key={data["id"]} data={data} />;
    });

    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.articles}
      >
        <Block flex center>
          {postsList}
        </Block>
    
      </ScrollView>
    );
  };

  render() {
    return (
      <Block flex center style={styles.home}>
        {this.renderPosts()}

      </Block>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,    
  },
  posts: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});

export default Friends;
